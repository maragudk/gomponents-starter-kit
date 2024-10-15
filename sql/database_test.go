package sql_test

import (
	"context"
	"testing"

	"maragu.dev/is"

	"app/sql"
)

func TestDatabase_GetThings(t *testing.T) {
	t.Run("gets things", func(t *testing.T) {
		db := sql.NewDatabase(sql.NewDatabaseOptions{})

		things, err := db.GetThings(context.Background())
		is.NotError(t, err)
		is.True(t, len(things) >= 2)
	})
}
