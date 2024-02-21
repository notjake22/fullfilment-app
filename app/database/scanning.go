package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ScanInItem(scanData ScanInRequest) (ScanInResponse, error) {
	payload, err := json.Marshal(scanData)
	if err != nil {
		return ScanInResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/scan/in", bytes.NewBuffer(payload))
	if err != nil {
		return ScanInResponse{}, err
	}

	req.Header = http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ScanInResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ScanInResponse{}, errors.New("item not found")
		}
		var response GeneralResponse
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return ScanInResponse{}, err
		}

		log.Println("Error sending scan item in: ", response.Error)
		return ScanInResponse{}, errors.New("error scanning item in")
	}

	var response ScanInResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ScanInResponse{}, err
	}

	return response, nil
}

func CancelScanItem(id string) error {
	payload, err := json.Marshal(map[string]interface{}{
		"itemID": id,
	})
	if err != nil {
		log.Println("1", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/scan/in/cancel", bytes.NewBuffer(payload))
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

func ScanOutItem(scanData ScanInRequest) (ScanInResponse, error) {
	payload, err := json.Marshal(scanData)
	if err != nil {
		return ScanInResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/scan/out", bytes.NewBuffer(payload))
	if err != nil {
		return ScanInResponse{}, err
	}

	req.Header = http.Header{
		"accept":       []string{"application/json"},
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ScanInResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			return ScanInResponse{}, errors.New("item not found")
		}
		var response GeneralResponse
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return ScanInResponse{}, err
		}

		log.Println("Error sending scan item in: ", response.Error)
		return ScanInResponse{}, errors.New("error scanning item in")
	}

	var response ScanInResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return ScanInResponse{}, err
	}

	return response, nil
}

func CancelScanOutItem(id string) error {
	payload, err := json.Marshal(map[string]interface{}{
		"itemID": id,
	})
	if err != nil {
		log.Println("1", err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/scan/out/cancel", bytes.NewBuffer(payload))
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
