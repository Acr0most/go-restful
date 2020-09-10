# about
This package provides functionality to generate rest APIs.

You can create your restful API with a given model and the matching database connection. There is no need for create the handling or write new routes and same code over and over again.

# features
- generate restful api endpoints with your custom handler
- generate restful api connected to database with gorm and default handler (json)

# working on
- add normal database connectors
- add multiple response format handlers like xml and csv
- bug fixing
- unit tests

# how-to
You don't want create multiple endpoints to create a RESTFUL api and add duplicated code?
Then this package could help you.

All what you need is a mapping from the url slug for an entity and a handler behind. 
Creating the routes and navigate the request to your custom handler is done by this package.   

# installation
```bash
go get github.com/Acr0most/go-restful
```
-- OR --
```go
import (
    "github.com/Acr0most/go-restful"
)
```

```bash
go mod tidy
```

# example
## gorm

- create your model

```go
import (
    // ...
    "github.com/Acr0most/go-restful/model"
)

type Example struct {
    model.CommonModelFields
    Name string `json:"name"`
}
```
- connect to database

```go
import (
    // ...
    "github.com/Acr0most/go-restful/connector"
)

connection := connector.NewGorm(connector.Config{MaxRetries: 10, IntervalMs: 1000})
connection.Connect(mysql.Open("<user>:<password>@tcp(<server/ip>:<port>)/<database>?charset=utf8mb4&parseTime=True&loc=Local"))
```

- init router
```go
import (
    // ...
    "github.com/Acr0most/go-restful/handler"
)

handle := handler.RestfulHandler{}

handle.InitRouter(handler.Config{
    "example": handler.HandlerConfig{
        Handler: handler.ConnectorHandler{Connector: connection},
        Dummy: handler.Dummy{
            Single:   &Example{},
            Multiple: &[]Example{},
        },
    },
}, 80)

err := handle.Handle()

if err != nil {
    panic(err)
}
```

## basic
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
