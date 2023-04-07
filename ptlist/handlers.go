package ptlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PtListGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		// De-serialize request
		params := &queryParams{}
		ok := params.fromRequestContext(c)
		if !ok {
			return
		}
		// Serialize response
		r := &ptListGetResponse{}
		r.fromTimestampsSlice(PeriodicTimestamps(params.period, params.t1, params.t2, params.tz))
		c.JSON(http.StatusOK, r.ptlist)
	}
}
