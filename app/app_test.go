package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ofortuna/app"
)

func TestMain(m *testing.M) int {

	os.Exit(m.Run())
}
// func TestGetFortuneReturnsOK(t *testing.T) {

// 	app.ListenAndServe(0)

	
		// req, _ := http.NewRequest(http.MethodGet, "/", nil)
		// resp := httptest.NewRecorder()

		// if resp.StatusCode != 200 {
		// 	t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		// }
		// got := resp.Body.String()
		// want := http.StatusOK

		// if got != want {
		// 	t.Errorf("got %q, want %q", got, want)
		// }
}
