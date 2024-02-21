package handlers

import (
	"github.com/gin-gonic/gin"
	"main/cache"
	"main/upc"
	"net/http"
)

func LookUpUPC(c *gin.Context, cRef cache.Cache) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid accept header",
			"success": false,
		})
		return
	}

	var upcCode string
	upcCode = c.Request.URL.Query().Get("upc")
	if upcCode == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "invalid upc param",
			"success": false,
		})
		return
	}

	u, err := upc.InitializeUPCLookup(upcCode, c.Request.Context(), cRef)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	itemData, err := u.RequestLookup()
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, map[string]interface{}{
				"error":   err,
				"success": false,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, ItemResponse{
		Name:     itemData.ItemName,
		ImageURI: itemData.ItemImageUri,
		Success:  true,
	})
}
