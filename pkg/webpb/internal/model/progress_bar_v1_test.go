package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgressBarV1_Inc_Updated(t *testing.T) {
	p := NewProgressBarV1(0, 10)
	assert.Equal(t, 0, p.Val())

	p.Inc()
	assert.Equal(t, 1, p.Val())
	assert.Equal(t, true, p.IsUpdated())
	assert.Equal(t, false, p.IsUpdated())
}

func TestProgressBarV1_Add_Updated(t *testing.T) {
	p := NewProgressBarV1(0, 10)
	assert.Equal(t, 0, p.Val())

	p.Add(0)
	assert.Equal(t, 0, p.Val())

	p.Add(1)
	assert.Equal(t, 1, p.Val())
	assert.Equal(t, true, p.IsUpdated())
	assert.Equal(t, false, p.IsUpdated())

	p.Add(20)
	assert.Equal(t, 10, p.Val())
	assert.Equal(t, true, p.IsUpdated())
	assert.Equal(t, false, p.IsUpdated())
}
