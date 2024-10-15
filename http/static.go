package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Static assets handler, which serves files from the root that have an extension, and everything from
// the images, scripts, and styles directories.
func Static(r chi.Router) {
	staticHandler := http.FileServer(http.Dir("public"))
	r.Get(`/{:[^.]+\.[^.]+}`, staticHandler.ServeHTTP)
	r.Get(`/{:images|scripts|styles}/*`, staticHandler.ServeHTTP)
}
