package handler

import (
	"encoding/json"
	"github.com/Acr0most/go-restful/connector"
	"net/http"
)

type ConnectorHandler struct {
	Connector connector.ConnectorInterface
}

func (t ConnectorHandler) Get(w http.ResponseWriter, r *http.Request) {
	dummy := r.Context().Value(KeyForConnectorPlaceholder)

	if success := t.Connector.Find(t.MapParamsFromQuery(r), dummy); !success {
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

	t.Connector.Delete(t.MapParamsFromQuery(r), dummy)
	_, _ = w.Write([]byte("example: delete"))
}

func (t ConnectorHandler) DeleteOne(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("example: delete one"))
}

func (t *ConnectorHandler) MapParamsFromQuery(r *http.Request) (params map[string]interface{}) {
	query := r.URL.Query()
	params = make(map[string]interface{}, len(query))

	for key, values := range query {
		params[key] = values
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
