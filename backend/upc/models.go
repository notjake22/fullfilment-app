package upc

import (
	"context"
)

type LookUp struct {
	UPC     string
	Context context.Context

	item ItemData
}

type ItemData struct {
	ItemName     string `json:"itemName"`
	ItemImageUri string `json:"itemImageUri"`
}

type ItemResponse struct {
	Name     string `json:"name"`
	ImageURI string `json:"imageURI"`
	Success  bool   `json:"success"`
}
