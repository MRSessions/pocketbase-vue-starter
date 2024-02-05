package main

import (
	"log"
	// "net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "pocket-base/migrations"

	// "github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: isGoRun, // auto creates migration files when making collection changes
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
		if getenvBool("POCKETBASE_DISABLE_UI") {
			e.Router.Pre(middleware.Rewrite(map[string]string{
				"/_":  "/",
				"/_*": "/",
			}))
			log.Default().Println("PocketBase UI is disabled")
		}
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDir), indexFallback))
		return nil
	})

	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	e.Router.GET("/api/init-check", func(c echo.Context) error {
	// 		var total int

	// 		err := app.DB().
	// 			Select("count(*)").
	// 			From("_admins").
	// 			Row(&total)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if total > 0 {
	// 			return c.JSON(http.StatusOK, map[string]interface{}{
	// 				"isSetup": true,
	// 				"message": "Initial setup is complete",
	// 			})
	// 		}
	// 		log.Default().Println("PocketBase admin user not found, initial setup needed")
	// 		return c.JSON(http.StatusOK, map[string]interface{}{
	// 			"isSetup": false,
	// 			"message": "PocketBase admin user not found, initial setup needed",
	// 		})
	// 	})
	// 	return nil
	// })

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

func getenvBool(key string) bool {
	val := os.Getenv(key)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	return ret
}
