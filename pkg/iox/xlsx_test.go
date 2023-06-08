package iox

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func cleanAfterTestsXlsx() (err error) {
	if err = os.Remove("test.xlsx"); err != nil {
		return
	}
	return
}

func TestWriterXlSX(t *testing.T) {
	defer func() {
		assert.NoError(t, cleanAfterTestsXlsx())
	}()

	w := NewWriterXLSX()
	defer func() {
		assert.NoError(t, w.Close())
	}()

	err := w.WriteRow("foo", RowXLSX("1", "2", 3), RowXLSX("1", "2", 3), RowXLSX("a", 2, 10))
	assert.NoError(t, err)

	err = w.WriteRow("foo", RowXLSX("1", "2", 3), RowXLSX("1", "2", 3), RowXLSX("a", 2, 10))
	assert.NoError(t, err)

	err = w.WriteRow("", RowXLSX("1", "2", 3), RowXLSX("1", "2", 3), RowXLSX("a", 2, 10))
	assert.NoError(t, err)

	err = w.WriteRow("bar", RowXLSX("1", "2", 3), RowXLSX("1", "2", 3), RowXLSX("a", 2, 10))
	assert.NoError(t, err)

	err = w.SaveAs("test.xlsx")
	assert.NoError(t, err)

}
