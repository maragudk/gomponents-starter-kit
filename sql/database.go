// Package sql provides a [Database].
package sql

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"

	"app/model"
)

// Database wraps a database connection pool.
type Database struct {
	log *slog.Logger
}

type NewDatabaseOptions struct {
	Log *slog.Logger
}

func NewDatabase(opts NewDatabaseOptions) *Database {
	if opts.Log == nil {
		opts.Log = slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	return &Database{
		log: opts.Log,
	}
}

// Connect to the database.
func (d *Database) Connect() error {
	d.log.Info("Connecting to database", "driver", "fake")
	return nil
}

// GetThings from the database.
func (d *Database) GetThings(ctx context.Context) ([]model.Thing, error) {
	var things []model.Thing
	n := rand.IntN(8) + 2

	for i := range n {
		things = append(things, model.Thing{Name: "Thing " + fmt.Sprint(i+1)})
	}

	return things, nil
}
