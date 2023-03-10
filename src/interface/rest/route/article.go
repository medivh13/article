package route

import (
	"net/http"

	handlers "article/src/interface/rest/handlers/article"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func ArticleRouter(h handlers.ArticleHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.Get("/", h.GetList)

	return r
}
