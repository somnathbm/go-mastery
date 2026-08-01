package scheduler

import (
	"context"
	"time"
)

type Scheduler struct {
	interval time.Duration
	callback func(context.Context)
}

func New(interval time.Duration, callback func(ctx context.Context)) *Scheduler {
	return &Scheduler{
		interval: interval,
		callback: callback,
	}
}

// recoverable callback
func (s *Scheduler) executeCallback(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			// logs panic
		}
	}()
	s.callback(ctx)
}

func (s *Scheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// as long as ticks are delivered on the channel, keep executing the callback
			s.executeCallback(ctx)
		case <-ctx.Done():
			// when Ctrl+C is pressed, it will gracefully shutdown by allowing in-flight callback to finish first
			return
		}
	}
}
