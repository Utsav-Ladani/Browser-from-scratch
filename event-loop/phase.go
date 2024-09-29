package main

type Phase interface {
	GetName() string
	IsEmpty() bool
	Run()
}
