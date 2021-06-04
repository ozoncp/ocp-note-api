package alarmer

import "time"

type Alarmer interface {
	Alarm() <-chan struct{}
}

type alarmer struct {
	duration time.Duration
	alarm    chan struct{}
	end      chan struct{}
}

func (a *alarmer) New(duration time.Duration) Alarmer {
	return &alarmer{
		duration: duration,
		alarm:    make(chan struct{}),
		end:      make(chan struct{}),
	}
}

func (a *alarmer) Init() {
	go func() {
		ticker := time.NewTicker(a.duration)
		defer ticker.Stop()
		defer close(a.alarm)

		for {
			select {
			case <-ticker.C:
				a.alarm <- struct{}{}
			case <-a.end:
				return
			}
		}
	}()
}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarm
}

func (a *alarmer) Close() {
	close(a.end)
}
