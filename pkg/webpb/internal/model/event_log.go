package model

func NewEventLog(len int) (ev EventLog) {
	return NewEventLogV1(len)
}
