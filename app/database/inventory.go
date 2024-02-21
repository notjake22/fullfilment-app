package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func GetAllItems() ([]ItemDataModel, error) {
	req, err := http.NewRequest(http.MethodGet, baseApiUrl+"/inventory/get-all", nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"accept": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var response GeneralResponse
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		log.Println("Error getting inventory items: ", response.Error)
		return nil, errors.New("error getting items")
	}

	var response InventoryResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New(response.Error)
	}

	return response.Items, nil
}

func DeleteIvenItem(itemId string) error {
	payload, err := json.Marshal(map[string]interface{}{
		"itemID": itemId,
	})
	if err != nil {
		log.Println("1", err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, baseApiUrl+"/inventory/delete-item", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("2", err)
		return err
	}

	req.Header = http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("3", err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		var response GeneralResponse
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			log.Println("4", err)
			return err
		}
		log.Println("Error sending cancel scan item in: ", response.Error)
		return errors.New("error scanning item in")
	}

	return nil
}
