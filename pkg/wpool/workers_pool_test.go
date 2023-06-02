package wpool

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestTask struct{}

func (t *TestTask) Do(_ context.Context) {
	time.Sleep(time.Second * 1)
	successCounter++
}

var successCounter = 0

func TestNewWPool(t *testing.T) {
	wp := NewWPool(context.Background(), 10).Start()

	clock := time.Now()
	var err error
	err = wp.Put(new(TestTask), new(TestTask), new(TestTask), new(TestTask), new(TestTask))
	assert.NoError(t, err)
	wp.Stop(false)

	check := time.Since(clock)

	assert.Equal(t, 5, successCounter)   // Все задачи выполнены
	assert.Less(t, check, time.Second*2) // Общее время выполнения должно быть меньше 2 секунд (будет 1с)
	// В 1 поток было бы минимум 5с.
	wp.Start()
	err = wp.Put(new(TestTask), new(TestTask), new(TestTask), new(TestTask), new(TestTask))
	assert.NoError(t, err)
	wp.Stop(false)

	assert.Equal(t, 10, successCounter)

	wp.Start()
	err = wp.Put(new(TestTask), new(TestTask), new(TestTask), new(TestTask), new(TestTask))
	assert.NoError(t, err)
	wp.Stop(true)

	assert.Less(t, successCounter, 15)
}
