package category_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/category"
	"github.com/gorilla/mux"
)

func TestCategoryDelete(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/category/delete/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "50"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(category.DeleteHandler)

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

func TestDeleteCategoryWrongID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/category/delete/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "ff"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(category.DeleteHandler)

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
