package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CheckBuildNumber(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid accept header",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"build": os.Getenv("CURRENT_BUILD"),
	})
}
