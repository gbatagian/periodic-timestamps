package settings

import "github.com/gin-gonic/gin"

var ENGINE *gin.Engine
var DatesLayout string

func init() {
	ENGINE = gin.Default()
	DatesLayout = "20060102T150405Z"
}
