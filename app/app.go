package app

import (
	"fmt"
	"os"
	"periodic-timestamps/settings"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func (app *App) Run() {
	app.Router.Run(
		fmt.Sprintf(
			"%s:%s",
			os.Getenv(settings.ApiHostEnvVarName),
			os.Getenv(settings.ApiPortEnvVarName),
		),
	)
}

func RunWithEngine(engine *gin.Engine) {
	app := App{Router: engine}
	app.GetRoutes()
	app.Run()
}
