package model

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"regexp"
)

type Statistics struct {
	Time   string
	Views  int
	Clicks int
	Cost   float32
	Cpc    float32
	Cpm    float32
}

func (s Statistics) Validate() error {

	return validation.ValidateStruct(&s,
		validation.Field(&s.Time, validation.Required, validation.Date("2006-01-02").Error("must be in format YYYY-MM-DD")),
		validation.Field(&s.Cost, validation.By(checkRuble)),
	)

}

func checkRuble(value interface{}) error {

	s := fmt.Sprintf("%.3f", value)
	return validation.Validate(s,
		validation.Match(regexp.MustCompile("^[0-9]*\\.[0-9]{2}0$")).Error("must be ruble"))
}
