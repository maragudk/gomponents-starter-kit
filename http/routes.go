package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"maragu.dev/httph"
)

// setupRoutes for the server.
func (s *Server) setupRoutes() {
	s.mux.Group(func(r chi.Router) {
		r.Use(middleware.Compress(5))

		// Sets up a static file handler with cache busting middleware.
		r.Group(func(r chi.Router) {
			r.Use(httph.VersionedAssets)

			Static(r)
		})

		Home(r, s.db)
	})
}
