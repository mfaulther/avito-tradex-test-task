package model

import (
	"encoding/json"
	"strings"
	"time"
)

type StatDate struct {
	Time time.Time
}

func (c *StatDate) UnmarshalJSON(b []byte) (err error) {

	const layout = "2006-01-02"

	s := strings.Trim(string(b), `"`)
	c.Time, err = time.Parse(layout, s)
	if err != nil {
		return err
	}
	return nil

}

func (c StatDate) MarshalJSON() ([]byte, error) {
	res, e := json.Marshal(c.Time.Format("2006-01-02"))
	if e != nil {
		return nil, e
	}
	return res, nil
}
