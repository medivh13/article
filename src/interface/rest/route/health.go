package route

import (
	"net/http"

	handlers "article/src/interface/rest/handlers/health"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func HealthRouter(h handlers.IHealthHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", h.Ping)
	r.Get("/health", h.HealthCheck)

	return r
}
