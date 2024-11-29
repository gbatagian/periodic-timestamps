package ptlist

import (
	"net/url"
	"regexp"
	"time"

	"periodic-timestamps/settings"
)

type QueryParams struct {
	Period Period
	Tz     *time.Location
	T1     time.Time
	T2     time.Time

	rawParams url.Values
}

func (qp *QueryParams) ParseURL(url *url.URL) (*QueryParams, error) {
	qp.rawParams = url.Query()

	if err := qp.parsePeriod(); err != nil {
		return nil, err
	}
	if err := qp.parseLocation(); err != nil {
		return nil, err
	}
	if err := qp.parseTimestamps(); err != nil {
		return nil, err
	}

	return qp, nil
}

func (qp *QueryParams) validateNotEmpty(name string) (string, error) {
	value := qp.rawParams.Get(name)

	if value == "" {
		return value, emptyQueryParameter{name: name}
	}

	return value, nil
}

func (qp *QueryParams) parsePeriod() error {
	value, err := qp.validateNotEmpty("period")
	if err != nil {
		return err
	}

	p, err := PeriodFromString(value)
	qp.Period = p

	return err
}

func (qp *QueryParams) parseLocation() error {
	value, err := qp.validateNotEmpty("tz")
	if err != nil {
		return err
	}

	tz, err := time.LoadLocation(value)
	qp.Tz = tz

	return err
}

func (qp *QueryParams) parseTimestamps() error {
	timestampAttrs := []string{"t1", "t2"}
	for _, name := range timestampAttrs {
		value, err := qp.validateNotEmpty(name)
		if err != nil {
			return err
		}

		pattern := `^\d{8}T\d{6}Z$`
		r := regexp.MustCompile(pattern)
		if !r.MatchString(value) {
			return invalidTimestampFormat{timestamp: value}
		}
	}

	layout := settings.DatesLayout

	t1, err := time.Parse(layout, qp.rawParams.Get("t1"))
	if err != nil {
		return err
	}
	t2, err := time.Parse(layout, qp.rawParams.Get("t2"))
	if err != nil {
		return err
	}

	qp.T1 = t1
	qp.T2 = t2

	return nil
}

type PtListGetResponse struct {
	Ptlist []string `json:"ptlist"`
}

func (r *PtListGetResponse) FromTimestampsSlice(tms []time.Time) *PtListGetResponse {
	var strTmsSlice []string
	for _, t := range tms {
		strTmsSlice = append(strTmsSlice, t.Format(settings.DatesLayout))
	}

	r.Ptlist = strTmsSlice
	return r
}
