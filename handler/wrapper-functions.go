package rest

import (
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	r.Context().Value(KeyForHandlerInterface).(HandlerInterface).Get(w, r)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
	r.Context().Value(KeyForHandlerInterface).(HandlerInterface).GetOne(w, r)
}

func Add(w http.ResponseWriter, r *http.Request) {
	r.Context().Value(KeyForHandlerInterface).(HandlerInterface).Add(w, r)
}

func AddOne(w http.ResponseWriter, r *http.Request) {
	r.Context().Value(KeyForHandlerInterface).(HandlerInterface).AddOne(w, r)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	r.Context().Value(KeyForHandlerInterface).(HandlerInterface).Delete(w, r)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	r.Context().Value(KeyForHandlerInterface).(HandlerInterface).DeleteOne(w, r)
}
