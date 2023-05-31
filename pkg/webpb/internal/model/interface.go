package model

type ProgressBar interface {
	Inc()
	Add(delta int)
	IsUpdated() bool
	Val() int
	Len() int
}

type EventLog interface {
	IsUpdated() (isUpdated bool, updates []string)
	Push(events ...string)
	Events() (events []string)
}
