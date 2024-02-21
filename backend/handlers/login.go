package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	if c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "invalid payload type",
			"success": false,
		})
		return
	}

	var request LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	val, ok := adminUsers[request.Username]
	if ok {
		if val != request.Password {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error":   "invalid user credentials",
				"success": false,
			})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "invalid user credentials",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"error":   "",
		"success": true,
	})
}
