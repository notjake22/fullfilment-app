package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"net/http"
	"sync"
)

func GetAllInventoryItems(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid accept header",
			"success": false,
		})
		return
	}

	items, err := database.GetAllInventoryItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, InventoryResponse{
		Items:   items,
		Success: true,
		Error:   "",
	})
}

func DeleteInventoryItem(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request ScanInCancelRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"error":   "",
		"success": true,
	})
}