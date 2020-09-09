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

func (t ExampleHandler) Patch(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: patch"))
}

func (t ExampleHandler) PatchOne(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: patch one"))
}
