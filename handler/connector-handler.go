package handler

import (
	"encoding/json"
	"github.com/Acr0most/go-restful/connector"
	"github.com/Acr0most/go-restful/middleware"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type ConnectorHandler struct {
	Connector connector.ConnectorInterface
}

func (t ConnectorHandler) Get(w http.ResponseWriter, r *http.Request) {
	dummy := r.Context().Value(KeyForConnectorPlaceholder)
	request := r.Context().Value(middleware.RequestKey).(middleware.Request)

	if err := t.Connector.Find(request.Filter, dummy); err != nil {
		w.WriteHeader(405)
		_, _ = w.Write([]byte("error while find entity: " + err.Error()))

		return

	}

	t.CreateResponse(w, dummy)
}

func (t ConnectorHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	t.Get(w, r)
}

func (t ConnectorHandler) Add(w http.ResponseWriter, r *http.Request) {
	dummy := r.Context().Value(KeyForConnectorPlaceholder)
	request := r.Context().Value(middleware.RequestKey).(middleware.Request)

	err := mapstructure.Decode(request.GetPayload(), &dummy)

	if err != nil {
		panic(err)
	}

	if err := t.Connector.Create(dummy); err != nil {
		w.WriteHeader(405)
		_, _ = w.Write([]byte("error while creating entity: " + err.Error()))

		return
	}

	t.CreateResponse(w, dummy)
}

func (t ConnectorHandler) AddOne(w http.ResponseWriter, r *http.Request) {
	t.Add(w, r)
}

func (t ConnectorHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (t ConnectorHandler) UpdateOne(w http.ResponseWriter, r *http.Request) {
	t.Update(w, r)
}

func (t ConnectorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	dummy := r.Context().Value(KeyForConnectorPlaceholder)
	request := r.Context().Value(middleware.RequestKey).(middleware.Request)

	if err := t.Connector.Delete(request.GetMerged(), dummy); err != nil {
		w.WriteHeader(405)
		_, _ = w.Write([]byte("error while delete entity: " + err.Error()))

		return
	}

	_, _ = w.Write([]byte("done"))
}

func (t ConnectorHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	t.Delete(w, r)
}

func (t ConnectorHandler) Patch(w http.ResponseWriter, r *http.Request) {
	dummy := r.Context().Value(KeyForConnectorPlaceholder)
	request := r.Context().Value(middleware.RequestKey).(middleware.Request)

	if err := t.Connector.Patch(request.Filter, request.GetPayload(), dummy); err != nil {
		w.WriteHeader(405)
		_, _ = w.Write([]byte("error while patching entity: " + err.Error()))

		return
	}

	_, _ = w.Write([]byte("todo.."))
}

func (t ConnectorHandler) PatchOne(w http.ResponseWriter, r *http.Request) {
	t.Patch(w, r)
}

func (t *ConnectorHandler) CreateResponse(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(object); err != nil {
		panic(err)
	}
}
