package main

type MicroTaskQueue struct {
	Queue[*Promise]
}

func (mtq *MicroTaskQueue) GetName() string {
	return "MicroTaskQueue"
}

func (mtq *MicroTaskQueue) IsEmpty() bool {
	return mtq.Length() == 0
}

func (mtq *MicroTaskQueue) Run() {
	pendingMicroTasksQueue := Queue[*Promise]{}

	for mtq.Length() > 0 {
		promise := mtq.Dequeue()

		if promise.status == PromiseStatusFulfilled {
			if promise.then != nil {
				promise.then.callback(promise.then.Resolve)
				pendingMicroTasksQueue.Enqueue(promise.then)
			}
		} else {
			pendingMicroTasksQueue.Enqueue(promise)
		}
	}

	mtq.Queue = pendingMicroTasksQueue
}

var MicroTask = &MicroTaskQueue{}

func (mtq *MicroTaskQueue) Promise(callback func(func())) *Promise {
	promise := &Promise{
		status:   PromiseStatusPending,
		callback: callback,
	}

	callback(promise.Resolve)
	println("")

	mtq.Enqueue(promise)

	return promise
}
