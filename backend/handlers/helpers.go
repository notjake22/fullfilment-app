package handlers

import (
	"context"
	"log"
	"main/database"
	"main/upc"
	"strings"
	"sync"
)

func checkForCustomItem(inUpc string, cancel context.CancelFunc, itemData *upc.ItemData, wg *sync.WaitGroup) {
	defer wg.Done()
	item, err := database.GetCustomItem(inUpc)
	if err != nil {
		log.Println("custom not found")
		return
	}
	itemData.ItemImageUri = item.ImageUri
	itemData.ItemName = item.ItemName
	cancel()
}

func lookupUpc(upcStr string, ctx context.Context, itemPointer *upc.ItemData, wg *sync.WaitGroup) error {
	defer wg.Done()
	u, err := upc.InitializeUPCLookup(upcStr, ctx)
	if err != nil {
		log.Println("Error init upc lookup: ", err)
		return err
	}
	item, err := u.RequestLookup()
	if err != nil {
		if err.Error() == "not found" || strings.Contains(err.Error(), "context canceled") {
			return nil
		}
		log.Println("Error looking up upc: ", err)
		return err
	}
	itemPointer.ItemName = item.ItemName
	itemPointer.ItemImageUri = item.ItemImageUri
	return nil
}
