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

func TestServer_Start(t *testing.T) {
	t.Run("can start and stop server", func(t *testing.T) {
		s := http.NewServer(http.NewServerOptions{})

		go func() {
			is.NotError(t, s.Start())
		}()
		defer func() {
			is.NotError(t, s.Stop())
		}()

		// I know we could check that the server is running here, but it's easier to just wait a bit
		time.Sleep(10 * time.Millisecond)

		res, err := http2.Get("http://localhost:8080/")
		is.NotError(t, err)
		is.Equal(t, http2.StatusOK, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		is.NotError(t, err)
		is.True(t, strings.HasPrefix(string(body), "<!doctype html>"))
	})
}
