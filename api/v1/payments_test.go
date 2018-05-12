package v1_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetPayments(t *testing.T) {
	URL := server.URL + "/api/v1/payments"
	res, err := http.Get(URL)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetPayment(t *testing.T) {
	URL := server.URL + "/api/v1/payments/1"
	res, err := http.Get(URL)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestPostPayment(t *testing.T) {
	URL := server.URL + "/api/v1/payments"
	res, err := http.PostForm(URL,
		url.Values{"placeid": {"7777"}, "cost": {"9999"}})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestUpdatePayment(t *testing.T) {
	URL := server.URL + "/api/v1/payments/1"
	jsonStr := `{"placeid": 7777, "cost": 9999}`

	req, err := http.NewRequest("PUT", URL, strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestDeletePayment(t *testing.T) {
	URL := server.URL + "/api/v1/payments/1"

	req, err := http.NewRequest("DELETE", URL, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}
