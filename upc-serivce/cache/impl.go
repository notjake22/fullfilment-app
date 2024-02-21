package cache

import (
	"context"
	"errors"
	"time"
)

func (c *CRef) Init() {
	c.data = map[string]ItemData{}
	go c.garbageCollect()
	return
}

func (c *CRef) Get(_ context.Context, value string) (ItemData, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.data[value]; !ok {
		return ItemData{}, errors.New("item does not exist")
	} else {
		c.data[value] = ItemData{
			TimeUsed: time.Now().Unix(),
			ImageUri: val.ImageUri,
			Upc:      val.Upc,
			ItemName: val.ItemName,
		}
		return val, nil
	}
}

func (c *CRef) Put(_ context.Context, value ItemData) {
	// redundant check to see if another instance already put same item in map before we attempt to
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[value.Upc] = value
}
