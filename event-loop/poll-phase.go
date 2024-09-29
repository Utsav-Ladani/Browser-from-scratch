package main

type PollPhase struct {
	queue Queue[IO]
}

func (pp *PollPhase) GetName() string {
	return "Poll"
}

func (pp *PollPhase) IsEmpty() bool {
	return pp.queue.Length() == 0
}

func (pp *PollPhase) Run() {
	maxIOOperations := 2

	for ; pp.queue.Length() > 0 && maxIOOperations > 0; maxIOOperations-- {
		io := pp.queue.Dequeue()
		io.Read()

		if io.status == IOStatusOpen {
			pp.queue.Enqueue(io)
		}
	}
}

var Poll = &PollPhase{}

func (pp *PollPhase) ReadFile(filename string, callback func(string, error)) {
	io := IO{
		readCallback: callback,
	}

	io.Open(filename)

	pp.queue.Enqueue(io)
}
