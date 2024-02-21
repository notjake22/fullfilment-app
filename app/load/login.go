package load

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Login(user string, pass string) (bool, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"username": user,
		"password": pass,
	})

	req, err := http.NewRequest(http.MethodPost, baseApiUrl+"/login", bytes.NewBuffer(payload))
	if err != nil {
		return false, err
	}

	req.Header = http.Header{
		"content-type": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusUnauthorized {
			return false, errors.New("invalid login")
		}
		return false, errors.New(fmt.Sprintf("request error checking login: %v", res.StatusCode))
	}

	var response LoginResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return false, err
	}

	return response.Success, nil
}
