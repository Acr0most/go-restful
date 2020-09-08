# about
This package provides functionality to generate rest routes mapped to your own handler.

# how-to
You don't want create multiple endpoints to create a RESTFUL api and add duplicated code?
Then this package helps you a lot.

All what you need is a mapping from the url slug for an entity and a handler behind. 
Creating the routes and navigate the request to your custom handler is done by this package.   

# example
In the example folder you can see a tiny project that creates the server for an example handler.

- Just create the restfulHandler
```go
handler := rest.RestfulHandler{}
```

- Set the port and mapping
```go
handler.InitRouter(map[string]rest.HandlerInterface{
    "example": ExampleHandler{},
}, 80)
```

- Start the server
```go
err := handler.Handle()
```

- handle errors
```go
if err != nil {
    panic(err)
}
```

Your handler have to match the HandlerInterface and provide this functions

```go
package main

import "net/http"

type ExampleHandler struct{}

func (t ExampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: get"))
}

func (t ExampleHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: get one"))
}

func (t ExampleHandler) Add(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: add"))
}

func (t ExampleHandler) AddOne(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: add one"))
}

func (t ExampleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: delete"))
}

func (t ExampleHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: delete one"))
}

```
