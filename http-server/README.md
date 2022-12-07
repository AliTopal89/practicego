#### HTTP Server

```go
func ListenAndServe(addr string, handler Handler) error
```
This will start a web server listening on a port, creating a goroutine for every request and running it against a `Handler`.

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.

Type implements the Handler interface by implementing the ServeHTTP method which expects two arguments, the first is where we write our response and the second is the HTTP request that was sent to the server.

