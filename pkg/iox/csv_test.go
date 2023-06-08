package iox

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCsvStruct struct {
	Col1 int       `csv:"col_1"`
	Col2 string    `csv:"col_2"`
	Col3 time.Time `csv:"col_3"`
}

func cleanAfterTestsCsv() (err error) {
	if err = os.Remove("test.csv"); err != nil {
		return
	}
	return
}

func TestWriterCSV(t *testing.T) {
	wr, err := NewWriterCSV("test.csv", TestCsvStruct{}, true)
	if err != nil {
		t.FailNow()
	}

	err = wr.WriteRow(testCsvRows)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, wr.Close())
	}()
}

func TestReaderCSV_ReadCSV(t *testing.T) {
	var rows []*TestCsvStruct

	err := NewReaderCSV().ReadCSV("test.csv", &rows)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(rows))

	assert.NoError(t, cleanAfterTestsCsv())
}

func TestUnmarshalCSV(t *testing.T) {
	var rows []*TestCsvStruct

	err := UnmarshalCSV([]byte(testCsvData), &rows)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(rows))
}

func TestMarshalCSV(t *testing.T) {
	data, err := MarshalCSV(testCsvRows)
	assert.NoError(t, err)
	assert.Equal(t, testCsvData, string(data))
}

const testCsvData = `col_1,col_2,col_3
123,1111,2023-06-04T01:01:01.000000001+03:00
124,2222,2023-06-04T02:01:01.000000001+03:00
`

var testCsvRows = []*TestCsvStruct{
	{
		Col1: 123,
		Col2: "1111",
		Col3: time.Date(2023, 06, 04, 1, 1, 1, 1, time.Local),
	},
	{
		Col1: 124,
		Col2: "2222",
		Col3: time.Date(2023, 06, 04, 1, 1, 1, 1, time.Local).Add(time.Hour),
	},
}
