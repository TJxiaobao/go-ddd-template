package vo

import (
	"strings"
	"time"
)

type Time struct {
	t time.Time
}

func NewTime(t time.Time) Time {
	return Time{t: t}
}

func NewTimeFromValue(t string) Time {
	date, _ := time.Parse("2006-01-02 15:04:05", t)
	return Time{t: date}
}

func NewDateTimeFromValue(t string) Time {
	t = strings.Split(t, " ")[0]
	date, _ := time.Parse("2006-01-02", t)
	return Time{t: date}
}

func (d Time) GetTime() time.Time {
	return d.t
}

func (d Time) GetDateValue() string {
	return d.t.Format("2006-01-02")
}
