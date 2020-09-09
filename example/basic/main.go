package main

import (
	"github.com/Acr0most/go-restful/handler"
)

type Example struct{}

func main() {
	handle := handler.RestfulHandler{}

	handle.InitRouter(handler.Config{
		Handler: map[string]handler.HandlerInterface{"example": ExampleHandler{}},
		Dummy: handler.Dummy{
			Single:   &Example{},
			Multiple: &[]Example{},
		},
	}, 80)

	err := handle.Handle()

	if err != nil {
		panic(err)
	}
}
