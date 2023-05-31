package model

import (
	"sync"
)

func NewProgressBarV1(val int, max int) *ProgressbarV1 {
	return &ProgressbarV1{
		val: val,
		len: max,
	}
}

type ProgressbarV1 struct {
	mu      sync.Mutex
	updated bool
	val     int
	len     int
}

func (p *ProgressbarV1) Inc() {
	if p.val < p.len {
		p.val++
		p.updated = true
	}
}

func (p *ProgressbarV1) Add(delta int) {
	if delta == 0 {
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if p.val+delta <= p.len {
		p.val += delta
		p.updated = true
	} else {
		if p.val != p.len {
			p.val = p.len
			p.updated = true
		}
	}
}

func (p *ProgressbarV1) IsUpdated() bool {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.updated {
		p.updated = false
		return true
	}
	return false
}

func (p *ProgressbarV1) Val() int {
	return p.val
}

func (p *ProgressbarV1) Len() int {
	return p.len
}
