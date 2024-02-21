package upc

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"main/cache"
	"net/http"
	"time"
)

func (u *LookUp) RequestLookup() (ItemData, error) {
	data, err := u.cache.Get(u.Context, u.UPC)
	if err != nil {
		log.Println("caching upc at end of req")
		u.shouldCache = true
	} else {
		return ItemData{
			ItemName:     data.ItemName,
			ItemImageUri: data.ImageUri,
		}, nil
	}

	req, err := http.NewRequestWithContext(u.Context, http.MethodGet, "https://www.barcodespider.com/"+u.UPC, nil)
	if err != nil {
		return ItemData{}, err
	}

	req.Header = http.Header{
		"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"Accept-Encoding": []string{"deflate"},
		"User-Agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15"},
		"Accept-Language": []string{"en-US,en;q=0.9"},
		"Connection":      []string{"keep-alive"},
		"Cache-Control":   []string{"no-cache"},
	}
	
	res, err := u.Client.Do(req)
	if err != nil {
		return ItemData{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Request status code: %v\n", res.StatusCode)
		if res.StatusCode == http.StatusNotFound {
			return ItemData{}, errors.New("not found")
		}
		if res.StatusCode == http.StatusTooManyRequests {
			// TODO: in another go routine, run this client to solve captcha for proxy
			// TODO: while in this one swap ips, and continue to run upc scan
			log.Println("Caught rate limit while scanning, retrying with proxy")
			err = u.proxyHandler()
			if err != nil {
				return ItemData{}, err
			}
			return u.RequestLookup()
		}
		return ItemData{}, errors.New(fmt.Sprintf("Request error getting item info: %v", res.StatusCode))
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return ItemData{}, err
	}

	err = u.findItemInfo(doc)
	if err != nil {
		return ItemData{}, err
	}

	if u.shouldCache {
		go u.cache.Put(u.Context, cache.ItemData{
			TimeUsed: time.Now().Unix(),
			ImageUri: u.item.ItemImageUri,
			Upc:      u.UPC,
			ItemName: u.item.ItemName,
		})
	}

	return u.item, nil
}
