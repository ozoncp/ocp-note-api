package alarmer

import (
	"time"
)

type Alarmer interface {
	Alarm() <-chan struct{}
	Init() error
	Close()
}

type alarmer struct {
	duration time.Duration
	alarm    chan struct{}
	end      chan struct{}
}

func New(duration time.Duration) Alarmer {

	if duration <= 0 {
		return nil
	}

	return &alarmer{
		duration: duration,
		alarm:    make(chan struct{}),
		end:      make(chan struct{}),
	}
}

func (a *alarmer) Init() error {
	go func() {
		ticker := time.NewTicker(a.duration)
		defer ticker.Stop()
		defer close(a.alarm)
		defer close(a.end)

		for {
			select {
			case <-ticker.C:
				select {
				case a.alarm <- struct{}{}:
				default:
				}
			case <-a.end:
				return
			}
		}
	}()

	return nil
}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarm
}

func (a *alarmer) Close() {
	a.end <- struct{}{}
}
