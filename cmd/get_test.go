package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
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
			wantOut: []byte(`HTTP/1.1: 200 OK
Content-Length: [14]
Content-Type: [text/plain; charset=utf-8]
`),
		},
	}

	for _, tt := range testCases {
		testServer := httptest.NewServer(tt.handler)
		defer testServer.Close()

		buf := bytes.NewBufferString("")
		if err := getURL(testServer.URL, buf); err != tt.wantErr {
			t.Errorf("Error not match: want %v but got %v", tt.wantErr, err)
		}

		got := ignoreLines(buf.Bytes(), "Date")

		if diff := cmp.Diff(tt.wantOut, got); diff != "" {
			t.Errorf("Output string differ (-want +got): \n%v", diff)
		}
	}
}
