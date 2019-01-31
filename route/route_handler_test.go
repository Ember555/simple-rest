package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestSearchParSuccess(t *testing.T) {
	router := mux.NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get("http://localhost:8080/searchPar")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /searchAll is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestSearchByNameSuccess(t *testing.T) {
	var d testData
	router := mux.NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get("http://localhost:8080/search/name/inintName")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(json.NewDecoder(res.Body).Decode(&d))

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /search/name/{name} is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func TestSearchByNameFail(t *testing.T) {
	router := mux.NewRouter()
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get("http://localhost:8080/search/name/inintName")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /search/name/{name} is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}
