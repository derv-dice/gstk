package interval

import (
	"fmt"
	"time"
)

type Interval struct {
	date1 time.Time
	date2 time.Time
}

const strLayout = "2006-01-02 15:04:05 .999999999"

func NewInterval(from time.Time, to time.Time) *Interval {
	return &Interval{date1: from, date2: to}
}

func (i *Interval) String() string {
	return i.Format(strLayout)
}

func (i *Interval) Format(layout string) string {
	x0, x1 := i.Sort()
	return fmt.Sprintf("%s | %s", x0.Format(layout), x1.Format(layout))
}

func (i *Interval) IsValid() bool {
	return !i.date1.IsZero() && !i.date2.IsZero()
}

// Sort - Возвращает даты Interval.date1 и Interval.date2 в хронологическом порядке
func (i *Interval) Sort() (d1, d2 time.Time) {
	if i.date2.After(i.date1) {
		d1 = i.date1
		d2 = i.date2
		return
	}

	d1 = i.date2
	d2 = i.date1
	return
}

// SortDesc - Возвращает даты Interval.date1 и Interval.date2 в хронологически обратном порядке
func (i *Interval) SortDesc() (d1, d2 time.Time) {
	if i.date2.After(i.date1) {
		d2 = i.date1
		d1 = i.date2
		return
	}

	d2 = i.date2
	d1 = i.date1
	return
}

func (i *Interval) Duration() time.Duration {
	x0, x1 := i.Sort()
	return x1.Sub(x0)
}

func (i *Interval) Divide(factor int, periodType PeriodType, precision time.Duration) (result []*Interval, err error) {
	if !i.IsValid() {
		err = fmt.Errorf("current Interval is not valid. Create new one with calling NewInterval()")
		return
	}

	discrete := true
	if periodType == Month || periodType == Year {
		discrete = false
	}

	xS, xF := i.Sort()
	var x0, x1 time.Time

	n := 1
	l := newPeriod(periodType, factor*n, discrete, xS).Duration()

	if l-precision <= 0 {
		err = fmt.Errorf("precission (%s) must be less than period (%s)", precision.String(), l.String())
		return
	}

	// Проверим, не оказалось ли так, что мы захотели разбить интервал на периоды, длительность которых больше
	// исходного интервала
	if l >= xF.Sub(xS) {
		result = append(result, i)
		return
	}

	for {
		x0 = x1
		if n == 1 {
			x0 = xS
		}

		l = newPeriod(periodType, factor, discrete, x0).Duration()
		x1 = x0.Add(l)

		if x1.After(xF) || x1 == xF {
			x1 = xF
			result = append(result, &Interval{date1: x0, date2: xF})
			break
		}

		result = append(result, &Interval{date1: x0, date2: x1.Add(-1 * precision)})
		n++
	}

	return
}
