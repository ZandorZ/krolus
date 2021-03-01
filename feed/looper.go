package feed

import (
	"time"
)

// Looper ...
/* loop := NewLooper(1*time.Second, func(tick time.Time) {
	fmt.Printf("\nHello %s", tick)
})
loop.Start(true)

time.Sleep(4 * time.Second)
loop.Pause()

fmt.Println("bla bla bla")

time.Sleep(2 * time.Second)
loop.Resume()

time.Sleep(4 * time.Second)
loop.Quit()
*/
type Looper struct {
	started  bool
	ticker   *time.Ticker
	interval time.Duration
	quit     chan bool
	task     func(time.Time)
}

// NewLooper ...
func NewLooper(interval time.Duration, task func(time.Time)) *Looper {
	return &Looper{
		ticker:   time.NewTicker(interval),
		interval: interval,
		task:     task,
	}
}

// Start ...
func (l *Looper) Start(now bool) {
	if !l.started {
		l.started = true
		l.quit = make(chan bool)
		if now {
			l.task(time.Now())
		}
		go l.routine()
	}
}

func (l *Looper) routine() {
	for {
		select {
		case <-l.quit:
			l.started = false
			return
		case tick := <-l.ticker.C:
			l.task(tick)
		}
	}
}

// Pause ...
func (l *Looper) Pause() {
	l.ticker.Stop()
}

// Resume ...
func (l *Looper) Resume() {
	l.ticker.Reset(l.interval)
}

// Quit ...
func (l *Looper) Quit() {
	close(l.quit)
}
