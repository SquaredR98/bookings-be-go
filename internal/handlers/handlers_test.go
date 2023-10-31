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

// TODO: Add test for the contact page
var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gs", "/general", "GET", []postData{}, http.StatusOK},
	{"ps", "/premium", "GET", []postData{}, http.StatusOK},
	{"sa", "/availability", "GET", []postData{}, http.StatusOK},
	{"reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-sa", "/availability", "POST", []postData{
		{key: "startDate", value: "2020-01-01"},
		{key: "endDate", value: "2020-01-01"},
	}, http.StatusOK},
	{"post-mr", "/make-reservation", "POST", []postData{
		{key: "name", value: "Ravi Ranjan"},
		{key: "email", value: "ravi@mail.com"},
		{key: "phone", value: "9098767876"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("For %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
