package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rizlantamima/dewantara/types"
)

func TestUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()

	HandleUsers(recorder, req)

	// Check the response status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check header
	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type %s but got %s", "application/json", contentType)
	}
	

	// Check the response body
	expected := []types.User{}
	err := json.Unmarshal(recorder.Body.Bytes(), &expected)
	if err != nil {
		t.Errorf("Failed to parse response JSON: %s", err)
	}}
