package job

import (
	"errors"
	"fmt"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type Scheduler struct {
	semaphore *semaphore.Weighted
	queue     queue
	running   int32
	closed    int32
}

func NewScheduler(concurrent int64) *Scheduler {
	return &Scheduler{
		semaphore: semaphore.NewWeighted(concurrent),
	}
}

func (s *Scheduler) Start() error {
	if err := s.checkClosed(); err != nil {
		return err
	}
	if !s.stopped() {
		return errors.New("already started")
	}
	atomic.StoreInt32(&s.running, 1)
	s.run()
	return nil
}

func (s *Scheduler) Stop() error {
	if err := s.checkClosed(); err != nil {
		return err
	}
	if s.stopped() {
		return errors.New("already stopped")
	}
	atomic.StoreInt32(&s.running, 0)
	return nil
}

func (s *Scheduler) Close() error {
	if err := s.checkClosed(); err != nil {
		return err
	}
	atomic.StoreInt32(&s.closed, 1)
	if err := s.Stop(); err != nil {
		return err
	}
	if length := s.queue.Len(); length > 0 {
		return fmt.Errorf("closed with non-empty queue: length %d", length)
	}
	return nil
}

func (s *Scheduler) Add(job Job) error {
	if err := s.checkClosed(); err != nil {
		return err
	}
	s.queue.Push(job)
	if !s.stopped() {
		s.run()
	}
	return nil
}

func (s *Scheduler) run() {
	if !s.semaphore.TryAcquire(1) {
		return
	}
	go func() {
		defer s.semaphore.Release(1)
		for !s.stopped() {
			run := s.queue.Pop()
			if run == nil {
				break
			}
			run()
		}
	}()
}

func (s *Scheduler) checkClosed() error {
	if atomic.LoadInt32(&s.closed) == 1 {
		return errors.New("scheduler closed")
	}
	return nil
}

func (s *Scheduler) stopped() bool {
	return atomic.LoadInt32(&s.running) == 0
}
