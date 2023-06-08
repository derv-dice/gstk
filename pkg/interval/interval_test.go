package interval

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const divider = "----------------------------------------------------------------------------------"

func TestInterval_Divide(t *testing.T) {
	now := time.Date(2022, 5, 7, 12, 47, 11, 0, time.UTC)
	interval := NewInterval(now, now.AddDate(5, 7, 0))

	intervals, err := interval.Divide(1, Year, time.Second)
	assert.NoError(t, err)
	fmt.Println(len(intervals))

	fmt.Printf("%s --> %s\n\n", now.Format(strLayout), now.Add(time.Hour*48).Format(strLayout))

	fmt.Println(divider)
	for i := range intervals {
		fmt.Println(intervals[i].String(), "|", intervals[i].Duration().String())
	}
	fmt.Println(divider)
	fmt.Println()
}
