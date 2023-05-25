package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWellcome(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	HandleWellcome(recorder, req)

	// Check the response status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello, World!"
	actualResponseString := recorder.Body.String()
	if actualResponseString != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actualResponseString, expected)
	}
}
