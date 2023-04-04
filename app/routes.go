package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckGet() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			map[string]string{
				"Felling": "Great",
			},
		)
	}
}

func (app *App) GetRoutes() {
	app.Router.GET("/healthcheck", HealthCheckGet())
}
