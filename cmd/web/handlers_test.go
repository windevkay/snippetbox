package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/windevkay/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	// Initialize a new httptest.ResponseRecorder. 
	rr := httptest.NewRecorder()
	// Initialize a new dummy http.Request.
	r, err := http.NewRequest(http.MethodGet, "/", nil) 
	if err != nil {
		t.Fatal(err) 
	}

	ping(rr, r)

	rs := rr.Result()
	// Check that the status code written by the ping handler was 200.
	assert.Equal(t, rs.StatusCode, http.StatusOK)
	// And we can check that the response body written by the ping handler // equals "OK".
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)

	if err != nil { 
		t.Fatal(err)
	} 

	bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK") 
}