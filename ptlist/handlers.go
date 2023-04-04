package ptlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PtListGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := &queryParams{}
		ok := params.fromRequestContext(c)
		if !ok {
			return
		}
		c.JSON(http.StatusOK, map[string]string{
			"period": params.period,
			"tz":     params.tz,
			"t1":     params.t1,
			"t2":     params.t2,
		})
	}
}
