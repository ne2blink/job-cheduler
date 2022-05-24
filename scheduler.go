package job

import (
	"context"
	"sync/atomic"

	"github.com/smallnest/queue"
	"golang.org/x/sync/semaphore"
)

type QueueScheduler struct {
	semaphore *semaphore.Weighted
	queue     *queue.CQueue
}

func NewQueueScheduler(concurrent int64) *QueueScheduler {
	return &QueueScheduler{
		semaphore: semaphore.NewWeighted(concurrent),
		queue:     queue.NewCQueue(),
	}
}

func (s *QueueScheduler) Add(job func()) {
	s.queue.Enqueue(job)
	s.run()
}

func (s *QueueScheduler) run() {
	if !s.semaphore.TryAcquire(1) {
		return
	}
	go func() {
		defer s.semaphore.Release(1)
		for {
			run := s.queue.Dequeue()
			if run == nil {
				return
			}
			run.(func())()
		}
	}()
}

type ChannelScheduler struct {
	context    context.Context
	concurrent int
	queue      chan func()
	running    int32
}

// NewChannelScheduler creates a new scheduler.
func NewChannelScheduler(ctx context.Context, concurrent int) *ChannelScheduler {
	return &ChannelScheduler{
		context:    ctx,
		concurrent: concurrent,
		queue:      make(chan func()),
	}
}

func (s *ChannelScheduler) SetQueueSize(size int) {
	s.queue = make(chan func(), size)
}

func (s *ChannelScheduler) Start() {
	if atomic.LoadInt32(&s.running) == 1 {
		return
	}
	atomic.StoreInt32(&s.running, 1)
	for i := 0; i < s.concurrent; i++ {
		go func() {
			for {
				select {
				case run := <-s.queue:
					run()
				case <-s.context.Done():
					return
				}
			}
		}()
	}
}

func (s *ChannelScheduler) Add(job func()) error {
	select {
	case s.queue <- job:
		return nil
	case <-s.context.Done():
		return context.Canceled
	}
}
