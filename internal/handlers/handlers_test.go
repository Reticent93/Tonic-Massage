package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
	{"post-search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2021-01-01"},
		{key: "end", value: ", 2021 - 01 - 02"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2021-01-01"},
		{key: "end", value: ", 2021 - 01 - 02"},
	}, http.StatusOK},
	{"booking", "/booking", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "jm@here.com"},
		{key: "phone", value: "555-123-1234"},
		{key: "date", value: "10-01-2021"},
		{key: "time", value: "09:00 a.m."},
	}, http.StatusOK},
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
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			res, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if res.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
			}
		}
	}
}
