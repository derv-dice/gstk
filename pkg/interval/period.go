package interval

import (
	"time"
)

type PeriodType int

const (
	Nanosecond PeriodType = iota + 1
	Microsecond
	Millisecond
	Second
	Minute
	Hour
	Day
	Week
	Month
	Year
)

type Period struct {
	periodType PeriodType
	duration   time.Duration
	startFrom  time.Time // Used only for Month and Year
}

func (p *Period) Duration() time.Duration {
	return p.duration
}

func NewPeriodNano(count int) *Period {
	return newPeriod(Nanosecond, count, true, time.Time{})
}

func NewPeriodMicro(count int) *Period {
	return newPeriod(Microsecond, count, true, time.Time{})
}

func NewPeriodMilli(count int) *Period {
	return newPeriod(Millisecond, count, true, time.Time{})
}

func NewPeriodSec(count int) *Period {
	return newPeriod(Millisecond, count, true, time.Time{})
}

func NewPeriodMin(count int) *Period {
	return newPeriod(Minute, count, true, time.Time{})
}

func NewPeriodHour(count int) *Period {
	return newPeriod(Hour, count, true, time.Time{})
}

func NewPeriodDay(count int) *Period {
	return newPeriod(Hour, count, true, time.Time{})
}

func NewPeriodWeek(count int) *Period {
	return newPeriod(Week, count, true, time.Time{})
}

func NewPeriodMonth(count int, startDate time.Time) *Period {
	return newPeriod(Month, count, false, startDate)
}

func NewPeriodYear(count int, startDate time.Time) *Period {
	return newPeriod(Year, count, false, startDate)
}

func newPeriod(kind PeriodType, count int, isDiscrete bool, startDate time.Time) (p *Period) {
	startDate = startDate.UTC()

	if count <= 0 {
		count = 1
	}

	p = &Period{
		periodType: kind,
	}

	switch kind {
	case Nanosecond:
		p.duration = time.Nanosecond
	case Microsecond:
		p.duration = time.Microsecond
	case Millisecond:
		p.duration = time.Millisecond
	case Second:
		p.duration = time.Second
	case Minute:
		p.duration = time.Minute
	case Hour:
		p.duration = time.Hour
	case Day:
		p.duration = time.Hour * 24
	case Week:
		p.duration = time.Hour * 168
	case Month:
		tmp := startDate.AddDate(0, count, 0)
		p.duration = tmp.Sub(startDate)
	case Year:
		tmp := startDate.AddDate(count, 0, 0)
		p.duration = tmp.Sub(startDate)
	}

	if isDiscrete {
		p.duration *= time.Duration(count)
	}

	return
}
