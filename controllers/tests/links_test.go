package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/torkashvand/goshortener/models"
)

func TestCreateLink(t *testing.T) {
	// Grab our router
	router, sqlite := initApp()
	defer sqlite.Close()

	randomAddress := "https://" + generateRandomString(8) + ".com"
	data := map[string]string{"address": randomAddress}

	w := performPostRequest(router, "/links", data)

	// Assert we encoded correctly,
	// the request gives a 201, w.Code
	if w.Code != http.StatusCreated {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", http.StatusCreated, w.Code)
	}
}

func TestGetLinks(t *testing.T) {
	// Grab our router and db connection
	router, sqlite := initApp()
	defer sqlite.Close()

	randomAddress := "https://" + generateRandomString(8) + ".com"
	link := models.Link{Address: randomAddress}

	sqlite.GetDB().Create(&link)

	// Perform a GET request with that handler.
	w := performGetRequest(router, "/links")

	// Assert we encoded correctly,
	// the request gives a 200, w.Code)
	if w.Code != http.StatusOK {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", http.StatusOK, w.Code)
	}

	// Convert the JSON response to a map
	var response map[string][]models.Link
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Grab the value & whether or not it exists
	// Make some assertions on the correctness of the response.
	if err != nil {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", nil, err)
	}

	count := len(response["data"])
	if count != 1 {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", 1, count)

	}

	returnedAddress := response["data"][0].Address
	if returnedAddress != randomAddress {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", randomAddress, returnedAddress)
	}
}

func TestRedirect(t *testing.T) {
	// Grab our router
	router, sqlite := initApp()
	defer sqlite.Close()

	randomAddress := "https://" + generateRandomString(8) + ".com"
	data := map[string]string{"address": randomAddress}

	w := performPostRequest(router, "/links", data)
	// Convert the JSON response to a map
	var response map[string]models.Link
	json.Unmarshal([]byte(w.Body.String()), &response)

	redirectAddress := response["data"].Shortcut
	res := performGetRequest(router, redirectAddress)

	// Assert we encoded correctly,
	// the request gives a 200, w.Code)
	if res.Code != http.StatusMovedPermanently {
		t.Errorf("TestConvertBase FAILED, expected value %v but got %v", http.StatusMovedPermanently, res.Code)
	}
}
