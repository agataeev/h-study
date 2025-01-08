package main

import (
	"context"
)

type Semaphore struct {
	ch chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, capacity)}
}

func (s *Semaphore) Acquire(ctx context.Context, n int64) error {
	for i := int64(0); i < n; i++ {
		select {
		case s.ch <- struct{}{}:
		case <-ctx.Done():
			// Освобождаем уже занятые места при прерывании
			for j := int64(0); j < i; j++ {
				<-s.ch
			}
			return ctx.Err()
		}
	}
	return nil
}

func (s *Semaphore) TryAcquire(n int64) bool {
	for i := int64(0); i < n; i++ {
		select {
		case s.ch <- struct{}{}:
		default:
			// Освобождаем уже занятые места
			for j := int64(0); j < i; j++ {
				<-s.ch
			}
			return false
		}
	}
	return true
}

func (s *Semaphore) Release(n int64) {
	for i := int64(0); i < n; i++ {
		<-s.ch
	}
}
