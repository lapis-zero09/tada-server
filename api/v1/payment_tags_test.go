package v1_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetPaymentTags(t *testing.T) {
	URL := server.URL + "/api/v1/payment_tags"
	res, err := http.Get(URL)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetPaymentTag(t *testing.T) {
	URL := server.URL + "/api/v1/payment_tags/1"
	res, err := http.Get(URL)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestPostPaymentTag(t *testing.T) {
	URL := server.URL + "/api/v1/payment_tags"
	res, err := http.PostForm(URL,
		url.Values{"paymentId": {"5"}, "tagId": {"12"}})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestUpdatePaymentTag(t *testing.T) {
	URL := server.URL + "/api/v1/payment_tags/1"
	jsonStr := `{"paymentId": 3, "tagId": 12}`

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

func TestDeletePaymentTag(t *testing.T) {
	URL := server.URL + "/api/v1/payment_tags/1"

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
