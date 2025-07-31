package api

import (
	"fmt"
	"github.com/bagasdisini/multifinance-api/docs"
	"github.com/bagasdisini/multifinance-api/internal/handler"
	"github.com/bagasdisini/multifinance-api/internal/pkg/config"
	"github.com/bagasdisini/multifinance-api/internal/pkg/log"
	_mysql "github.com/bagasdisini/multifinance-api/internal/pkg/mysql"
	"github.com/bagasdisini/multifinance-api/version"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"strings"
)

const appName = "Multifinance API"

func RunServer() {
	defer log.RecoverWithTrace()

	e := echo.New()
	log.SetLogger(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(config.CORSAllowOrigins, ","),
		AllowCredentials: true,
	}))

	apiBase(e)
	_mysql.RunMigration()

	handler.NewCustomerHandler(e, _mysql.DB)
	handler.NewTransactionHandler(e, _mysql.DB)

	log.Fatal(e.Start(fmt.Sprintf(`%v:%v`, config.AppHost, config.AppPort)))
}

func apiBase(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Multifinance API</title>
    </head>
    <body>
  		<h1>Welcome to %v</h1>
  		<p><a href="/api/version">version: %v</a></p>
  		<p><a href="/swagger/index.html#/">docs</a></p>
	</body>
</html>`, appName, version.Version))
	})

	e.GET("/api/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"name":    appName,
			"version": version.Version,
		})
	})

	docs.SwaggerInfo.Version = version.Version
	docs.SwaggerInfo.Host = config.SwaggerHost
	e.GET("/swagger/*", echoswagger.WrapHandler)
}
