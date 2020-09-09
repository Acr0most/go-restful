package handler

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"reflect"
	"strconv"
)

const KeyForHandlerInterface = "HANDLER_INTERFACE_CONTROL_KEY"
const KeyForConnectorPlaceholder = "CONNECTOR_PLACEHOLDER_CONTROL_KEY"

type Dummy struct {
	Single   interface{}
	Multiple interface{}
}

type Config struct {
	Handler map[string]HandlerInterface
	Dummy   Dummy
}

type RestfulHandler struct {
	Config Config
	Router *chi.Mux
	port   int
}

func (t *RestfulHandler) InitRouter(config Config, port int) {
	t.port = port
	t.Config = config

	t.Router = chi.NewRouter()
	t.Router.Use(middleware.Logger)

	t.Router.Route("/{config-element}/{id}", func(r chi.Router) {
		r.Use(t.AddContext)
		r.Get("/", GetOne)
		r.Post("/", AddOne)
		r.Delete("/", DeleteOne)
	})

	t.Router.Route("/{config-element}", func(r chi.Router) {
		r.Use(t.AddContext)
		r.Get("/", Get)
		r.Post("/", Add)
		r.Delete("/", Delete)
	})
}

func (t *RestfulHandler) AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context

		element := chi.URLParam(r, "config-element")
		isSingle := chi.URLParam(r, "id") != ""

		if _, exists := t.Config.Handler[element]; !exists {
			w.WriteHeader(404)
			_, _ = w.Write([]byte("Unknown requested entity. FIX by: adding >" + element + "< to an existing handler."))

			return
		}

		ctx = context.WithValue(r.Context(), KeyForHandlerInterface, t.Config.Handler[element])
		r.WithContext(ctx)

		switch isSingle {
		case true:
			ctx = context.WithValue(ctx, KeyForConnectorPlaceholder, reflect.New(reflect.ValueOf(t.Config.Dummy.Single).Elem().Type()).Interface())
		default:
			ctx = context.WithValue(ctx, KeyForConnectorPlaceholder, reflect.New(reflect.ValueOf(t.Config.Dummy.Multiple).Elem().Type()).Interface())
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *RestfulHandler) Handle() (err error) {
	err = http.ListenAndServe(":"+strconv.Itoa(t.port), t.Router)

	if err != nil {
		return err
	}

	return
}
