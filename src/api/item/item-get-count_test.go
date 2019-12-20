package item_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/item"
)

func TestGetCount(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/items/count", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.GetCountHandler)

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

	t.Run("response.body count", func(t *testing.T) {
		want := "53"
		if response := rr.Body.String(); strings.TrimSpace(response) != want {
			t.Errorf("Handler returned wrong count: got %v want %v", response, want)
		}
	})
}
