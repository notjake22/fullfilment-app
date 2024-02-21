package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/database"
	"net/http"
)

func GetServSettings(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid accept header",
			"success": false,
		})
		return
	}

	set, err := database.QueryAdminSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"error":     "",
		"success":   true,
		"itemPrice": set.ItemPrice,
	})
}

func PutServSettings(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request database.AdminSettings
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	err = database.SetNewAdminSettings(request)
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
