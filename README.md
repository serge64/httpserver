# httpserver

This package is a wrapper around the go standard HTTP server.

## Example of use

```go
func main() {
    mux := mux.NewServeMux()
    // some code
    srv := httpserver.New(&http.Server{Handler: mux})

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
    defer signal.Stop(quit)

    srv.Start()
    log.Println("server started")

    select {
    case <-quit:
        // closed database connections, etc.
    case err := <-srv.Notify():
        log.Fatalf("error %s", err)
    }

    if err := srv.Shutdown(context.Background()); err != nil {
        log.Fatalf("server graceful shutdown failed: %s", err)
    }

    log.Println("server graceful shutdown")
}
```
