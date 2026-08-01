package ratelimiter

import (
	"context"
	"sync"
	"time"
)

type RateLimiter struct {
	mu   sync.Mutex
	cond *sync.Cond

	capacity     int
	tokens       int
	refillAmount int
	refillEvery  time.Duration
}

func New(capacity int, refillAmount int, refillEvery time.Duration) *RateLimiter {
	rl := &RateLimiter{
		capacity:     capacity,
		tokens:       capacity,
		refillAmount: refillAmount,
		refillEvery:  refillEvery,
	}
	rl.cond = sync.NewCond(&rl.mu)
	return rl
}

func (r *RateLimiter) StartRefill(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(r.refillEvery)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				r.mu.Lock()
				r.tokens = min(r.capacity, r.tokens+r.refillAmount)
				r.mu.Unlock()
				r.cond.Broadcast()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (r *RateLimiter) Acquire() {
	r.mu.Lock()
	defer r.mu.Unlock()

	// if no tokens available, wait
	for r.tokens < 1 {
		r.cond.Wait()
	}
	// else tokens available
	r.tokens--
}
