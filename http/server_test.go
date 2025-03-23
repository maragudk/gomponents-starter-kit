package http_test

import (
	"io"
	http2 "net/http"
	"strings"
	"testing"
	"time"

	"maragu.dev/is"

	"app/http"
)

// waitForServer waits for the server up to one second to become ready.
func waitForServer(t *testing.T, url string) {
	client := &http2.Client{
		Timeout: 100 * time.Millisecond,
	}
	for i := 0; i < 50; i++ {
		resp, err := client.Get(url)
		if err == nil && resp.StatusCode == http2.StatusOK {
			resp.Body.Close()
			return
		}
		if err == nil {
			resp.Body.Close()
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("Server did not become ready in time")
}

func TestServer_Start(t *testing.T) {
	t.Run("can start and stop server", func(t *testing.T) {
		s := http.NewServer(http.NewServerOptions{})

		go func() {
			is.NotError(t, s.Start())
		}()
		defer func() {
			is.NotError(t, s.Stop())
		}()

		// on windows the server is not ready immediately
		waitForServer(t, "http://localhost:8080/")

		res, err := http2.Get("http://localhost:8080/")
		is.NotError(t, err)
		is.Equal(t, http2.StatusOK, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		is.NotError(t, err)
		is.True(t, strings.HasPrefix(string(body), "<!doctype html>"))
	})
}
