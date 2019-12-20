package item_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/item"
	"github.com/gorilla/mux"
)

func TestDelete(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/item/delete/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "50"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.DeleteHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("response.body", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})
}

func TestDeleteWrongID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/item/delete/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "ff"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.DeleteHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("response.body not null", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})
}
