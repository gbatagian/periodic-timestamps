package app

import (
	"net/http"

	"periodic-timestamps/ptlist"

	"github.com/gin-gonic/gin"
)

func healthCheckGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			map[string]string{
				"Felling": "Great",
			},
		)
	}
}

func (app *App) GetRoutes() {
	app.Router.GET("/healthcheck", healthCheckGet())
	app.Router.GET("/ptlist", ptlist.PtListGet())
}
