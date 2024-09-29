package main

type PromiseStatus string

const (
	PromiseStatusPending   PromiseStatus = "pending"
	PromiseStatusFulfilled PromiseStatus = "fulfilled"
)

type Promise struct {
	status   PromiseStatus
	callback func(func())
	then     *Promise
}

func (p *Promise) Resolve() {
	p.status = PromiseStatusFulfilled
}

func (p *Promise) Then(callback func(func())) *Promise {
	p.then = &Promise{
		status:   PromiseStatusPending,
		callback: callback,
	}

	if p.status == PromiseStatusFulfilled {
		callback(p.then.Resolve)
		println("")
	}

	return p.then
}
