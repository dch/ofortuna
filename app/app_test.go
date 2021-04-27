package app_test

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"ofortuna/app"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/phayes/freeport"
)

var serverURL string

func TestMain(m *testing.M) {
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	serverAddr := net.JoinHostPort("localhost", strconv.Itoa(port))
	go func() {
		err := app.ListenAndServe(serverAddr)
		if err != nil {
			log.Fatal(err)
		}
	}()
	conn, err := net.DialTimeout("tcp", serverAddr, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	serverURL = fmt.Sprintf("http://" + serverAddr + "/")
	os.Exit(m.Run())
}

func TestGetFortuneReturnsOK(t *testing.T) {
	resp, err := http.Get(serverURL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Errorf("got empty response body")
	}
}
