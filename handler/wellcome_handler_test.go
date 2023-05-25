package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWellcome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRec := httptest.NewRecorder()

	HandleWellcome(responseRec, req)

	// Check the response status code
	if status := responseRec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello, World!"
	actualResponseString := responseRec.Body.String()
	if actualResponseString != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actualResponseString, expected)
	}
}
