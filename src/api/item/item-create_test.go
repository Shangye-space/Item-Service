package item_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Shangye-space/Item-Service/src/api/item"
	"github.com/gorilla/mux"
)

func TestItemCreateRightData(t *testing.T) {
	expectedName := "TestItem"
	expectedPrice := 155.00
	expectedSubCategoryID := 5
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID": %v,"InSale":%v}`, expectedName, expectedPrice, expectedSubCategoryID, expectedInSale)

	req, err := http.NewRequest("POST", "/api/item/create", strings.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "50"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.CreateHandler)

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

func TestItemCreateWrongName(t *testing.T) {
	expectedPrice := 155.00
	expectedSubCategoryID := 5
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":5,"Price":%v,"SubCategoryID": %v,"InSale":%v}`, expectedPrice, expectedSubCategoryID, expectedInSale)

	req, err := http.NewRequest("POST", "/api/item/create", strings.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.CreateHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("response.body", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})
}

func TestItemCreateWrongPrice(t *testing.T) {
	expectedName := 555
	expectedSubCategoryID := 5
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":"%v","Price":-55,"SubCategoryID": %v,"InSale":%v}`, expectedName, expectedSubCategoryID, expectedInSale)

	req, err := http.NewRequest("POST", "/api/item/create", strings.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.CreateHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("response.body", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})
}
func TestItemCreateWrongSubCategoryID(t *testing.T) {
	expectedName := 555
	expectedPrice := 155.00
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID":-55,"InSale":%v}`, expectedName, expectedPrice, expectedInSale)

	req, err := http.NewRequest("POST", "/api/item/create", strings.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.CreateHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("response.body", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})
}

func TestItemCreateWrongInSale(t *testing.T) {
	expectedName := 555
	expectedPrice := 155.00
	expectedSubCategoryID := 5

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID": %v,"InSale":"rdfe"}`, expectedName, expectedPrice, expectedSubCategoryID)

	req, err := http.NewRequest("POST", "/api/item/create", strings.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(item.CreateHandler)

	handler.ServeHTTP(rr, req)

	t.Run("status", func(t *testing.T) {
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("response.body", func(t *testing.T) {
		if response := rr.Result().Body; response == nil {
			t.Errorf("Result was nil")
		}
	})
}
