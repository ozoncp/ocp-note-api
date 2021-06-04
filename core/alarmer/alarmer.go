package alarmer

import (
	"fmt"
	"time"
)

type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

type alarmer struct {
	duration time.Duration
	alarm    chan struct{}
	end      chan struct{}
}

func New(duration time.Duration) Alarmer {
	return &alarmer{
		duration: duration,
		alarm:    make(chan struct{}),
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
				fmt.Println("tik")
				select {
				case a.alarm <- struct{}{}:
					fmt.Println("tok")
				default:
					fmt.Println("tak")
				}
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
