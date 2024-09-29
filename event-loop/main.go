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

	fileContent := ""
	Poll.ReadFile("example.txt", func(data string, err error) {
		if err != nil {
			if err == ErrEof {
				println("File content: ", fileContent)
				println("")
			} else {
				println("Error: ", err.Error())
			}
		}

		fileContent += data
	})
}

func startEventLoop() {
	eventLoop := &EventLoop{
		microTasks: MicroTask,
	}
	eventLoop.AddPhase(Timer)
	eventLoop.AddPhase(Poll)
	eventLoop.AddPhase(Check)
	eventLoop.Run()
}
