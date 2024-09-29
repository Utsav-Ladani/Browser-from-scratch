package main

import "time"

type ScheduledCallback struct {
	callback func()
	time     time.Time
}

func (sc *ScheduledCallback) isScheduled() bool {
	return time.Now().After(sc.time)
}

func (sc *ScheduledCallback) run() {
	println("Running Timer callback")
	println("Time: ", time.Now().Format("2006-01-02 15:04:05"))
	println("Scheduled Time: ", sc.time.Format("2006-01-02 15:04:05"))
	sc.callback()
	println("")
}

type TimerPhase struct {
	queue Queue[ScheduledCallback]
}

func (tp *TimerPhase) GetName() string {
	return "Timer"
}

func (tp *TimerPhase) IsEmpty() bool {
	return tp.queue.Length() == 0
}

func (tp *TimerPhase) Run() {
	laterQueue := Queue[ScheduledCallback]{}

	for tp.queue.Length() > 0 {
		scheduledCallback := tp.queue.Dequeue()

		if scheduledCallback.isScheduled() {
			scheduledCallback.run()
		} else {
			laterQueue.Enqueue(scheduledCallback)
		}
	}

	tp.queue = laterQueue
}

var Timer = &TimerPhase{}

func (tp *TimerPhase) Timeout(callback func(), delay time.Duration) {
	tp.queue.Enqueue(ScheduledCallback{
		callback: callback,
		time:     time.Now().Add(delay),
	})
}
