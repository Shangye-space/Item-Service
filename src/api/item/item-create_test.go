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

func TestCreateRightData(t *testing.T) {
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

func TestCreateWrongName(t *testing.T) {
	var expectedName string
	expectedPrice := 155.00
	expectedSubCategoryID := 5
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID": %v,"InSale":%v}`, expectedName, expectedPrice, expectedSubCategoryID, expectedInSale)

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

func TestCreateWrongPrice(t *testing.T) {
	expectedName := 555
	var expectedPrice int
	expectedSubCategoryID := 5
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID": %v,"InSale":%v}`, expectedName, expectedPrice, expectedSubCategoryID, expectedInSale)

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
func TestCreateWrongSubCategoryID(t *testing.T) {
	expectedName := 555
	expectedPrice := 155.00
	var expectedSubCategoryID int
	expectedInSale := true

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID": %v,"InSale":%v}`, expectedName, expectedPrice, expectedSubCategoryID, expectedInSale)

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

func TestCreateWrongInSale(t *testing.T) {
	expectedName := 555
	expectedPrice := 155.00
	expectedSubCategoryID := 5
	var expectedInSale bool

	b := fmt.Sprintf(`{"Name":"%v","Price":%v,"SubCategoryID": %v,"InSale":%v}`, expectedName, expectedPrice, expectedSubCategoryID, expectedInSale)

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

