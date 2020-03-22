package idgen

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func TestListUsers(t *testing.T) {
	expected := Template()

	server := NewService()

	ts := httptest.NewServer(server)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := ts.Client().Do(req)
	if err != nil {
		t.Errorf("response error %q", err)
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)

	assertResponseBody(t, string(bodyBytes), expected)
}
