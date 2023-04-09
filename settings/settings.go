package settings

import "github.com/gin-gonic/gin"

var ENGINE *gin.Engine
var DatesLayout string
var ApiHostEnvVarName string
var ApiPortEnvVarName string

func init() {
	ENGINE = gin.Default()
	DatesLayout = "20060102T150405Z"
	ApiHostEnvVarName = "API_HOST"
	ApiPortEnvVarName = "API_PORT"
}
