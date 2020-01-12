package item_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/item"
	"github.com/gorilla/mux"
)

func TestGetByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/item/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.GetByIDHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("response.body not null", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})

	t.Run("response.body", func(t *testing.T) {
		expected := `[{"ID":1,"Name":"Jacket1","Price":125.99,"SubCategoryID":1,"InSale":true,"LastUpdated":null,"RemovedTime":null}]`
		response := rr.Body.String()
		actualRunes := []rune(strings.TrimSpace(response))
		actual := string(actualRunes[0:73]) + string(actualRunes[108:])
		if actual != expected {
			t.Errorf("Result was %v, but %v was expected", strings.TrimSpace(response), expected)
		}
	})
}
