package cache

import "time"

func (c *CRef) garbageCollect() {
	var gc = func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		for i, d := range c.data {
			if d.TimeUsed < time.Now().Unix()-(3600*3) {
				delete(c.data, i)
			}
		}
	}
	for {
		time.Sleep(time.Second * 5)
		gc()
	}
}
