package httpserver_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/serge64/httpserver"
)

func TestServer(t *testing.T) {
	srv := httpserver.New(&http.Server{})
	done := make(chan struct{})

	go srv.Start()

	go func() {
		t.Log("starting a graceful shutdown of the server")

		if err := srv.Shutdown(context.TODO()); err == nil {
			t.Log("server graceful shutdown")
		} else {
			t.Errorf("no error expected, but there is an error: %s", err)
		}

		close(done)
	}()

	err := <-srv.Notify()
	if err != http.ErrServerClosed {
		t.Errorf(
			"errors not equals:\n- expected: %s\n- actual: %s",
			http.ErrServerClosed,
			err,
		)
	}

	<-done
}
