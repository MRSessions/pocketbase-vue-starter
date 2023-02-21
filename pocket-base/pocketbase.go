package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "pocket-base/migrations"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	var publicDir string
	app.RootCmd.PersistentFlags().StringVar(
		&publicDir,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)

	var indexFallback bool
	app.RootCmd.PersistentFlags().BoolVar(
		&indexFallback,
		"indexFallback",
		true,
		"fallback the request to index.html on missing static path (eg. when pretty urls are used with SPA)",
	)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		var total int
		err := app.DB().
			Select("count(*)").
			From("_admins").
			Row(&total)
		if err != nil {
			return err
		}
		if total > 0 {
			e.Router.Pre(middleware.Rewrite(map[string]string{
				"/_":  "/",
				"/_/": "/",
			}))
			println("Initial Admin Already Setup")
		} else {
			e.Router.Pre(middleware.Rewrite(map[string]string{
				"/_":  "/",
				"/_*": "/",
			}))
			println()
			println("Initial Admin Setup Needed")
			println()
			println("PocketBase Admin Setup page ('/_/?installer#/') is replaced by the vue-client SetupAdmin page")
			println()
			println("After admin setup, you'll need to restart the container to access the PocketBase UI")
			println()
		}
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDir), indexFallback))
		return nil
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/api/init-check", func(c echo.Context) error {
			var total int

			err := app.DB().
				Select("count(*)").
				From("_admins").
				Row(&total)
			if err != nil {
				return err
			}
			if total > 0 {
				return c.JSON(http.StatusOK, map[string]interface{}{
					"isSetup": true,
					"message": "Initial setup is complete",
				})
			}
			return c.JSON(http.StatusOK, map[string]interface{}{
				"isSetup": false,
				"message": "Initial setup needed",
			})
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// the default pb_public dir location is relative to the executable
func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		return "./dist"
	}
	return filepath.Join(os.Args[0], "../dist")
}
