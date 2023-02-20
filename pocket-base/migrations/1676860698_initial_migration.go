// Left blank on purpose.
// You can use this file to add initial migrations
// or leave it blank and continue to add additional
// migrations. Automigrate is enabled

package migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
