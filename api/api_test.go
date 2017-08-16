package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetRoute(t *testing.T) {
	r := chi.NewRouter()
	addRoutes(r)
	ts := httptest.NewServer(r)
	defer ts.Close()
	rx := Response{}
	_, resp := testRequest(t, ts, "GET", "/api/id/1", nil)
	err := json.Unmarshal([]byte(resp), &rx)
	if err != nil {
		t.Fatalf("Response incorrect form ", resp)
	}
	if !reflect.DeepEqual(Response{Operation: "GET", Identifier: 1}, rx) {
    t.Fatalf("Response not equal expected values ", resp)
  }
}

func TestPostRoute(t *testing.T) {
	r := chi.NewRouter()
	addRoutes(r)
	ts := httptest.NewServer(r)
	defer ts.Close()
	rx := Response{}
	_, resp := testRequest(t, ts, "POST", "/api/id/1", nil)
	err := json.Unmarshal([]byte(resp), &rx)
	if err != nil {
		t.Fatalf("Response incorrect form ", resp)
	}
	if !reflect.DeepEqual(Response{Operation: "POST", Identifier: 1}, rx) {
    t.Fatalf("Response not equal expected values ", resp)
  }
}


//----------------------------------------------------------------------------
// HELPER METHOD
//----------------------------------------------------------------------------
func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (int, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return 0, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return resp.StatusCode, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return resp.StatusCode, ""
	}
	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)
}
