package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURL(t *testing.T) {
	testCases := []struct {
		handler http.HandlerFunc
		wantErr error
		wantOut []byte
	}{
		{
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "Hello, client")
			}),
			wantErr: nil,
		},
	}

	for _, tt := range testCases {
		testServer := httptest.NewServer(tt.handler)
		defer testServer.Close()

		buf := bytes.NewBufferString("")
		if err := getURL(testServer.URL, buf); err != tt.wantErr {
			t.Errorf("Error not match: want %v but got %v", tt.wantErr, err)
		}
	}
}
