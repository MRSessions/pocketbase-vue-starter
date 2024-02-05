// Left blank on purpose.
// You can use this file to add initial migrations
// or leave it blank and continue to add additional
// migrations. Automigrate is enabled

package migrations

import (
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		admin := &models.Admin{}
		adminEmail := getEnvVar("POCKETBASE_ADMIN_EMAIL", "admin@example.com")
		admin.Email = adminEmail
		adminPassword := getEnvVar("POCKETBASE_ADMIN_PASSWORD", "1234567890")
		admin.SetPassword(adminPassword)

		return dao.SaveAdmin(admin)
	}, func(db dbx.Builder) error { // optional revert operation

		dao := daos.New(db)

		adminEmail := getEnvVar("POCKETBASE_ADMIN_EMAIL", "admin@example.com")
		admin, _ := dao.FindAdminByEmail(adminEmail)
		if admin != nil {
			return dao.DeleteAdmin(admin)
		}

		// already deleted
		return nil
	})
}

func getEnvVar(name string, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultValue
	}
	return value
}
