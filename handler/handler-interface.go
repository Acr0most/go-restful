package handler

import "net/http"

type HandlerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
	AddOne(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeleteOne(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	PatchOne(w http.ResponseWriter, r *http.Request)
}
