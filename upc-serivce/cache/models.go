package cache

import (
	"context"
	"sync"
)

type Cache interface {
	// Get returns item data info about the upc
	// throws error if item does not exist
	Get(context.Context, string) (ItemData, error)

	// Put takes item data input and stores into cache
	Put(context.Context, ItemData)

	// Init starts and initializes cache map for use
	Init()
}

type CRef struct {
	mu   sync.Mutex
	data map[string]ItemData
}

type ItemData struct {
	TimeUsed int64  `json:"timeUsed"`
	ImageUri string `json:"imageUri"`
	ItemName string `json:"itemName"`
	Upc      string `json:"upc"`
}
