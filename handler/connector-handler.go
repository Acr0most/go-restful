package handler

import (
	"encoding/json"
	"github.com/Acr0most/go-restful/connector"
	"github.com/go-chi/chi"
	"net/http"
)

type ConnectorHandler struct {
	Connector connector.ConnectorInterface
}

func (t ConnectorHandler) Get(w http.ResponseWriter, r *http.Request) {
	dummy := r.Context().Value(KeyForConnectorPlaceholder)

	params := t.MapParamsFromQuery(r)
	id := chi.URLParam(r, "id")

	if id != "" {
		params["id"] = id
	}

	if success := t.Connector.Find(params, dummy); !success {
		dummy = nil
	}

	t.CreateResponse(w, dummy)
}

func (t ConnectorHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	t.Get(w, r)
}

func (t ConnectorHandler) Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	dummy := r.Context().Value(KeyForConnectorPlaceholder)
	err := decoder.Decode(dummy)

	if err != nil {
		panic(err)
	}

	t.Connector.Create(dummy)
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
	params := t.MapAllParams(r)
	id := chi.URLParam(r, "id")

	if id != "" {
		params["id"] = id
	}

	t.Connector.Delete(params, dummy)

	_, _ = w.Write([]byte("done"))
}

func (t ConnectorHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	t.Delete(w, r)
}

func (t *ConnectorHandler) MapParamsFromQuery(r *http.Request) (params map[string]interface{}) {
	query := r.URL.Query()
	params = make(map[string]interface{}, len(query))

	for key, values := range query {
		params[key] = values
	}

	return
}

func (t *ConnectorHandler) MapParamsFromPost(r *http.Request) (params map[string]interface{}) {
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		panic(err)
	}

	return
}

func (t *ConnectorHandler) MapAllParams(r *http.Request) (params map[string]interface{}) {
	params = t.MapParamsFromQuery(r)

	for key, value := range t.MapParamsFromPost(r) {
		params[key] = value // TODO: Merge?
	}

	return
}

func (t *ConnectorHandler) CreateResponse(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(object); err != nil {
		panic(err)
	}
}
