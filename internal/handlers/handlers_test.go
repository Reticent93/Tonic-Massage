package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/", "GET", []postData{}, http.StatusOK},
	{"mt", "/massage-therapists", "GET", []postData{}, http.StatusOK},
	{"lymph", "/lymphatic", "GET", []postData{}, http.StatusOK},
	{"sw", "/swedish", "GET", []postData{}, http.StatusOK},
	{"dt", "/deep-tissue", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"book", "/booking", "GET", []postData{}, http.StatusOK},
	{"ba", "/booking-summary", "GET", []postData{}, http.StatusOK},
}

func TestHandler(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			res, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
			}
		} else {

		}
	}
}
