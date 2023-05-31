package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventLogV1(t *testing.T) {
	ev := NewEventLogV1(50)
	assert.Equal(t, 50, len(ev.events))
}

func TestEventLogV1_Push_Updated(t *testing.T) {
	var isUpdated bool
	var updates []string

	ev := NewEventLogV1(4)
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, false, isUpdated)
	assert.Empty(t, updates)

	ev.Push("1", "2")
	assert.Equal(t, []string{"1", "2"}, ev.Events())
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, true, isUpdated)
	assert.Equal(t, []string{"1", "2"}, updates)
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, false, isUpdated)
	assert.Empty(t, updates)

	ev.Push("3", "4")
	assert.Equal(t, []string{"1", "2", "3", "4"}, ev.Events())
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, true, isUpdated)
	assert.Equal(t, []string{"3", "4"}, updates)
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, false, isUpdated)
	assert.Empty(t, updates)

	ev.Push("5")
	assert.Equal(t, []string{"2", "3", "4", "5"}, ev.Events())
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, true, isUpdated)
	assert.Equal(t, []string{"5"}, updates)
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, false, isUpdated)
	assert.Empty(t, updates)

	ev.Push("6", "7", "8", "9")
	assert.Equal(t, []string{"6", "7", "8", "9"}, ev.Events())
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, true, isUpdated)
	assert.Equal(t, []string{"6", "7", "8", "9"}, updates)
	isUpdated, updates = ev.IsUpdated()
	assert.Equal(t, false, isUpdated)
	assert.Empty(t, updates)
}
