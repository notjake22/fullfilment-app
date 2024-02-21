package database

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func GetServiceSettings(ctx context.Context) (ServiceSettings, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseApiUrl+"/service-settings", nil)
	if err != nil {
		return ServiceSettings{}, err
	}

	req.Header = http.Header{
		"accept": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ServiceSettings{}, err
	}

	if res.StatusCode != http.StatusOK {
		var response GeneralResponse
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return ServiceSettings{}, err
		}

		log.Println("Error getting service items: ", response.Error)
		return ServiceSettings{}, errors.New("error getting service items")
	}

	var response ServiceSettings
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return ServiceSettings{}, err
	}

	return response, nil
}

func SetNewServiceSettings(set ServiceSettings, ctx context.Context) error {
	payload, err := json.Marshal(set)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseApiUrl+"/service-settings/set-new", bytes.NewBuffer(payload))
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

	var response GeneralResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		log.Println("Error setting new service settings: ", response.Error)
		return errors.New("error setting new service settings")
	}

	if !response.Success {
		log.Println("Error setting new service settings: ", response.Error)
		return errors.New("error setting new service settings 2")
	}

	return nil
}
