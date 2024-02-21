package msHandlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PriceService int64

type ReqBody[C any] struct {
	Body C
}

const (
	AllPricing = iota
	SinglePricing
	NewPricing
	UpdatePricing
)

func GetAllPricing(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	response, err := requestPriceService(AllPricing, nil)
	if err != nil {
		log.Printf("Issue getting all prices from price service: %v\n", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var body []PricingPerItem
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "issue getting prices from service",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, body)
}

func GetPrice(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	response, err := requestPriceService(SinglePricing, c.Params.ByName("Id"))
	if err != nil {
		log.Printf("Issue getting all prices from price service: %v\n", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var body PricingPerItem
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "issue getting prices from service",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, body)
}

func NewPrice(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request PricingPerItem
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "invalid request body",
			"success": false,
		})
	}

	response, err := requestPriceService(NewPricing, request)
	if err != nil {
		log.Printf("Issue getting all prices from price service: %v\n", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var body PricingPerItem
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "issue getting prices from service",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, body)
}

func UpdatePrice(c *gin.Context) {
	if c.Request.Header.Get("accept") != "application/json" || c.Request.Header.Get("content-type") != "application/json" {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{
			"error":   "invalid type header",
			"success": false,
		})
		return
	}

	var request PricingPerItem
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "invalid request body",
			"success": false,
		})
	}

	response, err := requestPriceService(UpdatePricing, request)
	if err != nil {
		log.Printf("Issue getting all prices from price service: %v\n", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   err,
			"success": false,
		})
		return
	}

	var body PricingPerItem
	err = json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "issue getting prices from service",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, body)
}

func requestPriceService(service PriceService, body any) (*http.Response, error) {
	var req *http.Request
	var err error

	switch service {
	case AllPricing:
		req, err = http.NewRequest(http.MethodGet, "https://pricing-service-production.up.railway.app/pricing", nil)
		if err != nil {
			return nil, err
		}

		req.Header = http.Header{
			"accept": []string{"application/json"},
		}
	case SinglePricing:
		req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("https://pricing-service-production.up.railway.app/pricing/%v", body), nil)
		if err != nil {
			return nil, err
		}

		req.Header = http.Header{
			"accept": []string{"application/json"},
		}
	case NewPricing:
		var bytesBody []byte
		bytesBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(http.MethodPost, "https://pricing-service-production.up.railway.app/pricing/new", bytes.NewBuffer(bytesBody))
		if err != nil {
			return nil, err
		}

		req.Header = http.Header{
			"accept":       []string{"application/json"},
			"content-type": []string{"application/json"},
		}

	case UpdatePricing:
		var bytesBody []byte
		bytesBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(http.MethodPost, "https://pricing-service-production.up.railway.app/pricing/update", bytes.NewBuffer(bytesBody))
		if err != nil {
			return nil, err
		}

		req.Header = http.Header{
			"accept":       []string{"application/json"},
			"content-type": []string{"application/json"},
		}

	default:
		return nil, errors.New("invalid enum")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Request error, status code: %v", res.StatusCode))
	}
	return res, err
}
