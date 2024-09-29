package main

import "time"

func main() {
	defer startEventLoop()

	Timer.Timeout(func() {
		println("Timeout 1")
	}, 1*time.Second)

	Check.Immediate(func() {
		println("Immediate 1")
	})

	MicroTask.Promise(func(resolve func()) {
		Timer.Timeout(func() {
			resolve()
			println("Promise 1")
		}, 2*time.Second)
	}).Then(func(resolve func()) {

		Check.Immediate(func() {
			println("Immediate 2")
			resolve()
			println("Promise 1 then")
		})
	})
}

func startEventLoop() {
	eventLoop := &EventLoop{
		microTaskQueue: MicroTask,
	}
	eventLoop.AddPhase(Timer)
	eventLoop.AddPhase(Check)
	eventLoop.Run()
}
