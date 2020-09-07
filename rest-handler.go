package rest

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

const KeyForHandlerInterface = "HANDLER_INTERFACE_ACCESS_KEY"

type RestfulHandler struct {
	Config map[string]HandlerInterface
	router *chi.Mux
}

func (t *RestfulHandler) InitRouter(config map[string]HandlerInterface) {
	t.Config = config

	t.router = chi.NewRouter()
	t.router.Use(middleware.Logger)

	t.router.Route("/single/{element}", func(r chi.Router) {
		r.Use(t.AddContext)
		r.Get("/", GetOne)
		r.Post("/", AddOne)
		r.Delete("/", DeleteOne)
	})

	t.router.Route("/{element}", func(r chi.Router) {
		r.Use(t.AddContext)
		r.Get("/", Get)
		r.Post("/", Add)
		r.Delete("/", Delete)
	})

	return
}

func (t *RestfulHandler) AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		element := chi.URLParam(r, "element")

		if _, exists := t.Config[element]; !exists {
			w.WriteHeader(404)
			_, _ = w.Write([]byte("Unknown requested entity. FIX by: adding >" + element + "< to an existing handler."))

			return
		}

		ctx := context.WithValue(r.Context(), KeyForHandlerInterface, t.Config[element])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *RestfulHandler) Handle() (err error) {
	err = http.ListenAndServe(":80", t.router)

	if err != nil {
		return err
	}

	return
}
