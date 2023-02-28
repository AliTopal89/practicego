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

Type implements the `Handler` interface by implementing the `ServeHTTP` method which expects two arguments, the first is where we write our response and the second is the HTTP request that was sent to the server.

The `HandlerFunc` type is an adapter to allow the use of ordinary functions as HTTP handlers. If `f` is a function with the appropriate signature, `HandlerFunc(f)` is a Handler that calls `f`.

```go
type HandlerFunc func(ResponseWriter, *Request)
```

```go
server := http.HandlerFunc(PlayerServer)
log.Fatal(http.ListenAndServe(":5000", server))
```

From the documentation, we see that type `HandlerFunc` has already implemented the `ServeHTTP` method. By type casting our PlayerServer function with it, we have now implemented the required Handler.

`ListenAndServe` takes a port to listen on a `Handler`. If there is a problem the web server will return an error, an example of that might be the port already being listened to. For that reason we wrap the call in 	`log.Fatal` to log the error to the user.


```go
switch r.Method {...}
//
```
field `Method` string
Method specifies the HTTP method (GET, POST, PUT, etc.). For client requests, an empty string means GET

**Notes on type casting/conversion:**

- Type conversion happens when we assign the value of one data type to another
- What is the need for Type Conversion? 
  - Well, if you need to take advantage of certain characteristics of data type hierarchies, then we have to change entities from one data type into another. The general syntax for converting a value `val` to a `type T` is `T(val)`. 

- ```go
   var geek1 int = 845
  // explicit type conversion
  var geek2 float64 = float64(geek1)
  var geek3 int64 = int64(geek1)
  var geek4 uint = uint(geek1)
  ``` 
- Syntax of Type Conversion in go:
   - `newDataTypeVariable = T(oldDataTypeVariable)`, `T` is the new data type.

- ```go
   // taking the required
    // data into variables
    var x int = 19
    var y int = 21
    var mul float32

    // explicit type conversion
    mul = float32(x) * float32(y)

	// Displaying the result
    fmt.Printf("Multiplication = %f\n", mul)

	// Output: 
	// Multiplication = 399.000000
  ```
- if you don't to explicit type conversion:
  - ```go
      // taking the required
      // data into variables
      var x int = 19
      var y float32 = 21
      var mul float32

      mul = x * y

      // Displaying the result
      fmt.Printf("Multiplication = %f\n", mul)

	  // Output: invalid operation: x * y (mismatched types int and float32)
    ```

`ServeMux` is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

Patterns name fixed, rooted paths, like `"/favicon.ico"`, or rooted subtrees, like `"/images/"` (note the trailing slash). Longer patterns take precedence over shorter ones, so that if there are handlers registered for both `"/images/"` and `"/images/thumbnails/"`, the latter handler will be called for paths beginning `"/images/thumbnails/"` and the former will receive requests for any other paths in the `"/images/"` subtree.

Package `httprouter` is a trie based high performance HTTP request router.

```go
package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
```

The router matches incoming requests by the request method and the path. If a handle is registered for this path and method, the router delegates the request to that function. For the methods `GET, POST, PUT, PATCH and DELETE` shortcut functions exist to register handles, for all other methods router.Handle can be used.


**General Reminder Notes**:
`&` - variable's memory address
`*` - holds a memory address and resolves it, goes and gets the thing that the pointer was pointing at