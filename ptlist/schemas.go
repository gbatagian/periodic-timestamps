package ptlist

import (
	"fmt"
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

func (p *queryParams) validate(c *gin.Context) bool {
	return (p.validateParametersAreNotMissing(c) &&
		p.validateTimestampsHaveSpecifiedFormat(c) &&
		p.validatePeriod())
}

func (p *queryParams) validateParametersAreNotMissing(c *gin.Context) bool {
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

func (p *queryParams) validateTimestampsHaveSpecifiedFormat(c *gin.Context) bool {
	pattern := `^\d{8}T\d{6}Z$`
	r := regexp.MustCompile(pattern)
	for _, t := range []string{p.t1, p.t2} {
		if !r.MatchString(t) {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": fmt.Sprintf(
						"Timestamp '%s' does not match the required format, YYYYDDMMTHHMMSSZ", t,
					),
				},
			)
			return false
		}
	}
	return true
}

func (p *queryParams) validatePeriod() bool {
	p.period = PeriodFromString(p.period).String()
	return true
}

func (p *queryParams) fromRequestContext(c *gin.Context) bool {
	defer func() {
		// All unknown errors while parsing the URL query parameters we result in 400
		if r := recover(); r != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": r,
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
