package main

import (
	"github.com/Acr0most/go-restful"
)

func main() {
	handler := rest.RestfulHandler{}

	handler.InitRouter(map[string]rest.HandlerInterface{
		"example": ExampleHandler{},
	}, 80)

	err := handler.Handle()

	if err != nil {
		panic(err)
	}
}
