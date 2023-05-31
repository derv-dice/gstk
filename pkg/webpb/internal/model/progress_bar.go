package model

func NewProgressBar(val int, max int) ProgressBar {
	return NewProgressBarV1(val, max)
}
