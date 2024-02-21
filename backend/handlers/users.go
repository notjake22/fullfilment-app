package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"net/http"
	"sync"
)

func GetUsers(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid accept header",
			"success": false,
		})
		return
	}

	users, err := database.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var response UsersResponse
	response = UsersResponse{
		Success: true,
		Error:   "",
		Users:   nil,
	}
	var wg sync.WaitGroup
	wg.Add(len(users.Users))
	for _, user := range users.Users {
		go func(u database.User) {
			defer wg.Done()
			items, err := database.GetUserItems(u.ID)
			if err != nil {
				log.Println(err)
			}

			curStored := 0
			for _, item := range items {
				if item.CurrentlyStored == true {
					curStored++
				}
			}

			response.Users = append(response.Users, User{
				ID:                   u.ID,
				Name:                 u.Name,
				Phone:                u.Phone,
				Balance:              u.Balance,
				CurrentlyStoredItems: curStored,
				TotalOrderHistory:    len(items),
			})
		}(user)
	}
	wg.Wait()

	c.JSON(http.StatusOK, response)
}

func SetNewClient(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request database.User
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	err = database.SetNewUser(request)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
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

func EditClient(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request database.User
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	err = database.EditUser(request)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
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
