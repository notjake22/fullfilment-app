package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func GetAllUsers() ([]User, error) {
	req, err := http.NewRequest(http.MethodGet, baseApiUrl+"/user-list", nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Accept": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request error getting users list, status code: %v", res.StatusCode))
	}

	var response UsersResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New(response.Error)
	}

	return response.Users, nil
}

func SetNewClient(user User) bool {
	payload, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return false
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/add-new-client", bytes.NewBuffer(payload))
	if err != nil {
		return false
	}

	req.Header = http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Request error setting new user: %v\n", res.StatusCode)
		return false
	}

	var response GeneralResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return false
	}

	if !response.Success {
		return false
	}

	return true
}

func EditClient(user User) bool {
	payload, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return false
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/edit-client", bytes.NewBuffer(payload))
	if err != nil {
		return false
	}

	req.Header = http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Request error setting new user: %v\n", res.StatusCode)
		return false
	}

	var response GeneralResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return false
	}

	if !response.Success {
		return false
	}

	return true
}
