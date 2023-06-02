package wpool

import (
	"context"
	"errors"
	"sync"
)

type Task interface {
	Do(ctx context.Context)
}

// WPool - workers pool
type WPool struct {
	parentCtx context.Context
	ctx       context.Context

	count int
	in    chan Task
	ins   []chan Task
	out   chan Task
	stop  context.CancelFunc
	wg    sync.WaitGroup

	enabled bool
	stopped bool
}

func NewWPool(ctx context.Context, count int) *WPool {
	if count <= 0 {
		count = 1
	}

	childCtx, cancel := context.WithCancel(ctx)

	wp := &WPool{
		parentCtx: ctx,
		ctx:       childCtx,
		stop:      cancel,
		count:     count,
		in:        make(chan Task),
		ins:       []chan Task{},
		out:       make(chan Task),
	}

	for i := 0; i < count; i++ {
		wp.ins = append(wp.ins, make(chan Task))
	}

	return wp
}

// Put - Добавление задачи в пул
func (w *WPool) Put(tasks ...Task) (err error) {
	if !w.enabled {
		if w.count == 0 {
			return errors.New("can't put task into non created WPool. Create this one with NewWPool")
		}
		return errors.New("can't put task into stopped WPool. Use WPool.Start() before put new task into stopped WPool")
	}

	for i := range tasks {
		w.in <- tasks[i]
		w.wg.Add(1)
	}

	return
}

// Start - Включение работы пула задач.
// После этого можно использовать WPool.Put(t Task) чтобы добавлять задачи на выполнение
func (w *WPool) Start() *WPool {
	if w.enabled {
		return nil
	}

	if w.stopped {
		*w = *NewWPool(w.parentCtx, w.count)
	}

	// Распределение задач между каналами
	go func() {
		for {
			for _, c := range w.ins {
				select {
				case <-w.ctx.Done():
					return
				case task := <-w.in:
					c <- task
				}
			}
		}
	}()

	// Обработка задач из каналов
	for i := 0; i < w.count; i++ {
		go func(ctx context.Context, i int) {
			for {
				select {
				case <-ctx.Done():
					return
				case task := <-w.ins[i]:
					cCtx, cancel := context.WithCancel(w.ctx)
					task.Do(cCtx)
					w.wg.Done()
					cancel()
				}
			}
		}(w.ctx, i)
	}

	w.enabled = true
	return w
}

// Stop - Ждет, пока оставшиеся задачи закончат свое выполнение и выключает работу пула.
// При force=true незавершенные задачи отбрасываются и пул сразу выключается
func (w *WPool) Stop(force bool) {
	if !w.enabled || w.stopped {
		return
	}

	if !force {
		w.wg.Wait()
	}

	w.stop()
	w.enabled = false
	w.stopped = true
}
