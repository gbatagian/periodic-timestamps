package ptlist

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type queryParams struct {
	period string
	tz     string
	t1     string
	t2     string
}

func (p *queryParams) validate(c *gin.Context) bool {
	paramsMap := map[string]string{
		"period": p.period,
		"tz":     p.tz,
		"t1":     p.t1,
		"t2":     p.t2,
	}
	for k, v := range paramsMap {
		if v == "" {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": fmt.Sprintf(
						"The '%s' parameter cannot be empty. Please provide a value for this parameter.", k,
					),
				},
			)
			return false
		}
	}
	return true
}

func (p *queryParams) fromRequestContext(c *gin.Context) bool {
	period := c.Query("period")
	tz := c.Query("tz")
	t1 := c.Query("t1")
	t2 := c.Query("t2")
	p.period = period
	p.tz = tz
	p.t1 = t1
	p.t2 = t2
	return p.validate(c)
}
