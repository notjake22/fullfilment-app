package upc

import (
	"context"
	"main/cache"
	"net/http"
)

type LookUp struct {
	UPC     string
	Client  http.Client
	Context context.Context

	item        ItemData
	shouldCache bool

	cache cache.Cache
}

type ItemData struct {
	ItemName     string `json:"itemName"`
	ItemImageUri string `json:"itemImageUri"`
}
