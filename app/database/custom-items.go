package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func SetNewCustomItem(item CustomUPCModel) error {
	payload, err := json.Marshal(item)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/scan/new-item", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header = http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	var response GeneralResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Issue setting new custom item, status code: %v| reason: %s", res.StatusCode, response.Error))
	}

	if !response.Success {
		return errors.New("issue setting custom item: " + response.Error)
	}

	return nil
}
