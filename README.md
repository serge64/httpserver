# httpserver

This package is a wrapper around the go standard HTTP server.

## Example of use

```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "github.com/serge64/httpserver"
)

func main() {
    if err := run(); err != {
        log.Fatal(err)
    }
}

func run() error {
    mux := mux.NewServeMux()
    // ...
    server := http.Server{Handler: mux}
    srv := httpserver.New(&server)

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    defer signal.Stop(quit)

    go srv.Start()

    log.Println("server started")

    select {
    case <-quit:
        log.Println("\nsignal received")
        // closing the database connection, etc.
    case err := <-srv.Notify():
        return err
    }

    if err := srv.Shutdown(context.Background()); err != nil {
        log.Fatalf("server graceful shutdown failed: %s", err)
    }

    log.Println("server graceful shutdown")
    return nil
}
```
