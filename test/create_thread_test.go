package test

import (
	"bytes"
	"encoding/json"
	"goproject/go-site-backend/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateThread(t *testing.T) {

	var jsonStr = []byte(`{"title":"test assignment title", "content": "test assignment content", "userId":"8"}`)
	req, _ := http.NewRequest("POST", "/homework/postAssignment", bytes.NewBuffer(jsonStr))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.PostThread)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	if m["Title"] != "test thread" {
		t.Errorf("Expected product name to be 'test thread'. Got '%v'", m["Title"])
	}

	if m["Content"] != "test content" {
		t.Errorf("Expected product price to be 'test content'. Got '%v'", m["Content"])
	}

	if m["UserId"] != "8" {
		t.Errorf("Expected product ID to be '8'. Got '%v'", m["UserId"])

	}
}
