package main

type CheckPhase struct {
	queue Queue[func()]
}

func (cp *CheckPhase) GetName() string {
	return "Check"
}

func (cp *CheckPhase) IsEmpty() bool {
	return cp.queue.Length() == 0
}

func (cp *CheckPhase) Run() {
	for cp.queue.Length() > 0 {
		callback := cp.queue.Dequeue()

		println("Running check callback")
		callback()
		println("")
	}
}

var Check = &CheckPhase{}

func (cp *CheckPhase) Immediate(callback func()) {
	cp.queue.Enqueue(callback)
}
