package ptlist

import (
	"net/http"
	"periodic-timestamps/settings"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type queryParams struct {
	period  Period
	tz      *time.Location
	t1      time.Time
	t2      time.Time
	sPeriod string
	sTz     string
	sT1     string
	sT2     string
}

func (qp *queryParams) validate(c *gin.Context) bool {
	return (qp.validateParametersAreNotMissing(c) &&
		qp.validateTimestampsHaveSpecifiedFormat(c) &&
		qp.validatePeriod(c) &&
		qp.validateTimeZone(c))
}

func (qp *queryParams) validateParametersAreNotMissing(c *gin.Context) bool {
	paramsMap := map[string]string{
		"period": qp.sPeriod,
		"tz":     qp.sTz,
		"t1":     qp.sT1,
		"t2":     qp.sT2,
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
	for _, t := range []string{qp.sT1, qp.sT2} {
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
	layout := settings.DatesLayout

	qp.t1, _ = time.Parse(layout, qp.sT1)
	qp.t2, _ = time.Parse(layout, qp.sT2)
	return true
}

func (qp *queryParams) validatePeriod(c *gin.Context) bool {
	p, err := PeriodFromString(qp.sPeriod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"desc":   err.Error(),
		})
		return false
	}
	qp.period = p
	return true
}

func (qp *queryParams) validateTimeZone(c *gin.Context) bool {
	tz, err := time.LoadLocation(qp.sTz)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "error",
				"desc":   err.Error(),
			},
		)
		return false
	}
	qp.tz = tz
	return true
}

func (qp *queryParams) fromRequestContext(c *gin.Context) bool {
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

	qp.sPeriod = c.Query("period")
	qp.sTz = c.Query("tz")
	qp.sT1 = c.Query("t1")
	qp.sT2 = c.Query("t2")
	return qp.validate(c)
}

type ptListGetResponse struct {
	ptlist []string
}

func (r *ptListGetResponse) fromTimestampsSlice(tms []time.Time) {
	var strTmsSlice []string
	for _, t := range tms {
		strTmsSlice = append(strTmsSlice, t.Format(settings.DatesLayout))
	}

	r.ptlist = strTmsSlice
}
