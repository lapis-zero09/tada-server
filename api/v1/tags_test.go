package v1_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetTags(t *testing.T) {
	URL := server.URL + "/api/v1/tags"
	res, err := http.Get(URL)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetTag(t *testing.T) {
	URL := server.URL + "/api/v1/tags/1"
	res, err := http.Get(URL)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestPostTag(t *testing.T) {
	URL := server.URL + "/api/v1/tags"
	res, err := http.PostForm(URL,
		url.Values{"tagName": {"test tag"}})
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code is %d, got %d", http.StatusOK, res.StatusCode)
	}
}

func TestUpdateTag(t *testing.T) {
	URL := server.URL + "/api/v1/tags/1"
	jsonStr := `{"tagName":"updated tag"}`

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

func TestDeleteTag(t *testing.T) {
	URL := server.URL + "/api/v1/tags/1"

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
