package io

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func cleanAfterTestsTxt() (err error) {
	if err = os.Remove("test.txt"); err != nil {
		return
	}

	return
}

func TestReadFileLineByLine(t *testing.T) {
	defer func() {
		assert.NoError(t, cleanAfterTestsTxt())
	}()

	err := os.WriteFile("test.txt", []byte(testTxtFileData), 0644)
	assert.NoError(t, err)

	lines, err := ReadFileLineByLine("test.txt", true)
	assert.NoError(t, err)

	assert.Equal(t, 5, len(lines))
}

const testTxtFileData = `
foo
bar


Foo gg
     
Bar

baz`
