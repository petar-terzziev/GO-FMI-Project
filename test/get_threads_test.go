package test

import (
	"goproject/go-bank-backend/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetThreads(t *testing.T) {
	req, err := http.NewRequest("GET", "/forum/getAll", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetThreads)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ID":1,"Title":"Tema 1","Content":"Purva sedmica","UserId":"1"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
