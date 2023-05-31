package model

import (
	"sync"
)

func NewEventLogV1(len int) (ev *EventLogV1) {
	return &EventLogV1{
		events: make([]string, len, len),
		len:    len,
	}
}

type EventLogV1 struct {
	mu        sync.Mutex
	isUpdated bool
	events    []string
	head      int
	len       int
	updates   []string
}

func (e *EventLogV1) Push(events ...string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for i := range events {
		if e.head == e.len {
			e.head = 0
		}

		e.events[e.head] = events[i]
		e.head += 1
		e.isUpdated = true
		e.updates = append(e.updates, events[i])
	}
}

func (e *EventLogV1) IsUpdated() (isUpdated bool, updates []string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.isUpdated {
		isUpdated = true
		updates = e.updates
		e.isUpdated = false
		e.updates = []string{}
		return
	}

	return
}

func (e *EventLogV1) Events() (res []string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for i := range e.events[e.head:] {
		if e.events[e.head:][i] != "" {
			res = append(res, e.events[e.head:][i])
		}
	}

	for i := range e.events[:e.head] {
		if e.events[:e.head][i] != "" {
			res = append(res, e.events[:e.head][i])
		}
	}

	return
}
