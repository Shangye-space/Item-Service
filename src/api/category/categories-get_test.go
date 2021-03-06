package category_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/category"
)

func TestCategoriesGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/categories", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(category.GetHandler)

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
