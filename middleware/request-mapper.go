package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const (
	RequestKey = "KEY_FOR_REQUEST_OBJECT"
)

type Request struct {
	Filter   map[string]interface{}
	Payload  []map[string]interface{}
	Id       string
	IsSingle bool
}

func (t *Request) GetPayload() interface{} {
	if len(t.Payload) == 0 {
		return t.Payload
	}

	if t.IsSingle {
		return t.Payload[0]
	}

	return t.Payload
}
func (t *Request) GetMerged() (params map[string]interface{}) {
	params = t.Filter

	for _, el := range t.Payload {
		for key, value := range el {
			params[key] = value // TODO: override or merge?
		}
	}

	return
}

func RequestMapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filter := mapParamsFromQuery(r)
		payload, isSingle := mapParamsFromPost(r)
		id := chi.URLParam(r, "id")

		if id != "" {
			filter["id"] = id
		}

		ctx := context.WithValue(r.Context(), RequestKey, Request{Filter: filter, Payload: payload, Id: id, IsSingle: isSingle})
		r.WithContext(ctx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func mapParamsFromQuery(r *http.Request) (params map[string]interface{}) {
	query := r.URL.Query()
	params = make(map[string]interface{}, len(query))

	for key, values := range query {
		params[key] = values
	}

	return
}

func mapParamsFromPost(r *http.Request) (params []map[string]interface{}, isSingle bool) {
	if r.Method == "GET" {
		return
	}

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(r.Body)
	str := buf.String()

	if len(str) == 0 {
		return
	}

	switch str[:1] {
	case "[":
		isSingle = false
		params = mapSliceOfInterface(str)
	default:
		isSingle = true
		params = mapInterfaceMap(str)
	}

	return
}

func mapSliceOfInterface(str string) (params []map[string]interface{}) {
	err := json.Unmarshal([]byte(str), &params)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return
}

func mapInterfaceMap(str string) (params []map[string]interface{}) {
	var pseudo map[string]interface{}

	err := json.Unmarshal([]byte(str), &pseudo)

	if err != nil {
		log.Println(err.Error())
		return
	}

	return []map[string]interface{}{pseudo}
}
