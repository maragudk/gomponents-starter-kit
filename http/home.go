package http

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx/http"
	ghttp "maragu.dev/gomponents/http"

	"app/html"
	"app/model"
)

type thingsGetter interface {
	GetThings(ctx context.Context) ([]model.Thing, error)
}

// Home handler for the home page, as well as HTMX partial for getting things.
func Home(r chi.Router, db thingsGetter) {
	r.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		things, err := db.GetThings(r.Context())
		if err != nil {
			return nil, err
		}

		if hx.IsRequest(r.Header) {
			return html.ThingsPartial(things, time.Now()), nil
		}

		return html.HomePage(html.PageProps{}, things, time.Now()), nil
	}))
}
