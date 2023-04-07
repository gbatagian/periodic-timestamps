package ptlist

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type queryParams struct {
	period string
	tz     string
	t1     string
	t2     string
}

func (qp *queryParams) validate(c *gin.Context) bool {
	return (qp.validateParametersAreNotMissing(c) &&
		qp.validateTimestampsHaveSpecifiedFormat(c) &&
		qp.validatePeriod(c))
}

func (qp *queryParams) validateParametersAreNotMissing(c *gin.Context) bool {
	paramsMap := map[string]string{
		"period": qp.period,
		"tz":     qp.tz,
		"t1":     qp.t1,
		"t2":     qp.t2,
	}
	for k, v := range paramsMap {
		if v == "" {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": "error",
					"desc":   emptyQueryParameter{name: k}.Error(),
				},
			)
			return false
		}
	}
	return true
}

func (qp *queryParams) validateTimestampsHaveSpecifiedFormat(c *gin.Context) bool {
	pattern := `^\d{8}T\d{6}Z$`
	r := regexp.MustCompile(pattern)
	for _, t := range []string{qp.t1, qp.t2} {
		if !r.MatchString(t) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": "error",
					"desc":   invalidTimestampFormat{timestamp: t}.Error(),
				},
			)
			return false
		}
	}
	return true
}

func (qp *queryParams) validatePeriod(c *gin.Context) bool {
	_, err := PeriodFromString(qp.period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"desc":   err.Error(),
		})
		return false
	}
	return true
}

func (p *queryParams) fromRequestContext(c *gin.Context) bool {
	defer func() {
		// All unknown errors while parsing the URL query parameters we result in 400
		if r := recover(); r != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"status": "error",
					"desc":   r,
				},
			)
		}
	}()

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
