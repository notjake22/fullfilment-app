package upc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (u *LookUp) RequestLookup() (ItemData, error) {
	req, err := http.NewRequestWithContext(u.Context, http.MethodGet, "http://103.27.77.54:9900/upc-lookup?upc="+u.UPC, nil)
	if err != nil {
		return ItemData{}, err
	}

	req.Header = http.Header{
		"Accept":          []string{"application/json"},
		"Accept-Language": []string{"en-US,en;q=0.9"},
		"Connection":      []string{"keep-alive"},
		"Cache-Control":   []string{"no-cache"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ItemData{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ItemData{}, errors.New("not found")
		}

		return ItemData{}, errors.New(fmt.Sprintf("Request error getting item info: %v", res.StatusCode))
	}

	var response ItemResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ItemData{}, err
	}
	
	u.item = ItemData{
		ItemName:     response.Name,
		ItemImageUri: response.ImageURI,
	}

	return u.item, nil
}
