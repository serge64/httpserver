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

	srv.Start()

	go func() {
		t.Log("starting a graceful shutdown on the server")

		err := srv.Shutdown(context.TODO())
		if err == nil {
			t.Log("server graceful shutdown")
		} else {
			t.Errorf("no expected error but there is an error: %s", err)
		}

		close(done)
	}()

	err := <-srv.Notify()
	if err != http.ErrServerClosed {
		t.Error("expected an error server closed but received nil")
	}

	<-done
}
