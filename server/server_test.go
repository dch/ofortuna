package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/dch/ofortuna/server"
)

func TestGetFortune(t *testing.T) {
	t.Run("returns 200 OK and some text", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()

		FortuneServer(resp, req)

		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}
		got := resp.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
