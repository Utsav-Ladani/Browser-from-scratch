package main

type EventLoop struct {
	phases         []*Phase
	microTasks *MicroTaskQueue
}

func (el *EventLoop) AddPhase(phase Phase) {
	el.phases = append(el.phases, &phase)
}

func (el *EventLoop) Run() {
	for !el.isAllPhasesEmpty() {
		el.runEachPhase()
	}
}

func (el *EventLoop) isAllPhasesEmpty() bool {
	for _, phase := range el.phases {
		if !(*phase).IsEmpty() {
			return false
		}
	}

	return true
}

func (el *EventLoop) runEachPhase() {
	for _, phase := range el.phases {
		(*phase).Run()
		el.microTasks.Run()
	}
}
