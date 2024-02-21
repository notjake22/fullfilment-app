package load

import (
	"encoding/json"
	"log"
	"net/http"
)

// CheckBuild if func returns -1 there is no update, else it will return new version
func CheckBuild() int {
	req, err := http.NewRequest(http.MethodGet, baseApiUrl+"/build/check-version", nil)
	if err != nil {
		return -1
	}

	req.Header = http.Header{
		"accept": []string{"application/json"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return -1
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Error getting build num: %v\n", res.StatusCode)
		return -1
	}

	var response BuildResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return -1
	}

	if response.BuildNumber != buildNumber {
		return response.BuildNumber
	} else {
		return -1
	}
}
