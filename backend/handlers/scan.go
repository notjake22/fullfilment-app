package handlers

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"main/upc"
	"net/http"
	"sync"
)

// ScanIn
// Takes in request from client app with upc and client ID to scan into
// validates that user is in db and then looks up upc from reversed upc db
// gets admin price data, in 2 go routines, edits user balance, and adds item to db
func ScanIn(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request ScanRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	user, err := database.GetUserByID(request.ClientID)
	if err != nil {
		log.Println("Error getting user info: ", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var item upc.ItemData
	var wg1 sync.WaitGroup
	wg1.Add(2)
	go func() {
		ctx, cancelFunc := context.WithCancel(context.Background())
		go checkForCustomItem(request.UPC, cancelFunc, &item, &wg1)
		go func() {
			err = lookupUpc(request.UPC, ctx, &item, &wg1)
		}()
	}()
	wg1.Wait()

	if item.ItemImageUri == "" && item.ItemName == "" {
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error":   err,
				"success": false,
			})
			return
		}
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "item not found",
			"success": false,
		})
		return
	}

	settings, err := database.QueryAdminSettings()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = database.EditUser(database.User{
			ID:      user.ID,
			Name:    user.Name,
			Phone:   user.Phone,
			Balance: user.Balance + settings.ItemPrice,
		})
		if err != nil {
			log.Println("Error setting user new balance", err)
		}
	}()

	var itemInventoryID string
	go func() {
		defer wg.Done()
		itemInventoryID, err = database.SetNewItem(database.ItemDataModel{
			ItemName:        item.ItemName,
			ImageUri:        item.ItemImageUri,
			UPC:             request.UPC,
			OwnerName:       user.Name,
			OwnerID:         user.ID,
			CurrentlyStored: true,
		})
		if err != nil {
			log.Println("Error storing new item: ", err)
			return
		}
	}()
	wg.Wait()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, ScanInResponse{
		ID:       itemInventoryID,
		ItemName: item.ItemName,
		ImageUri: item.ItemImageUri,
	})
}

func ScanInCancel(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request ScanInCancelRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	itemData, err := database.GetItemByID(request.ItemID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		user, err := database.GetUserByID(itemData.OwnerID)
		if err != nil {
			log.Println("Error getting user data from id stored with item data: ", err)
			return
		}

		settings, err := database.QueryAdminSettings()
		if err != nil {
			log.Println(err)
			return
		}

		err = database.EditUser(database.User{
			ID:      user.ID,
			Name:    user.Name,
			Phone:   user.Phone,
			Balance: user.Balance - settings.ItemPrice,
		})
		if err != nil {
			log.Println("Error editing user balance: ", err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = database.DeleteItem(request.ItemID)
		if err != nil {
			log.Println("Error deleting item: ", err)
			return
		}
	}()
	wg.Wait()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}
}

func ScanOut(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request ScanRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	item, err := database.CheckUPCByOwner(request.ClientID, request.UPC)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	if item.ID == "" {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error":   "item not found with user",
			"success": false,
		})
		return
	}

	err = database.MarkItemOut(item.ID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, ScanInResponse{
		ID:       item.ID,
		ItemName: item.ItemName,
		ImageUri: item.ImageUri,
	})
}

func ScanOutCancel(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request ScanInCancelRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	err = database.MarkItemOut(request.ItemID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"error":   "",
		"success": true,
	})
}