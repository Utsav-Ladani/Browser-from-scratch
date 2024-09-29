package main

import (
	"errors"
	"math/rand"
)

const fileContent = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."

var ErrEof = errors.New("eof")

type IOStatus string

const (
	IOStatusOpen   IOStatus = "open"
	IOStatusClosed IOStatus = "closed"
	IOStatusError  IOStatus = "error"
)

type IO struct {
	status          IOStatus
	filename        string
	totalChars      int
	currentPosition int
	readCallback    func(string, error)
}

func (io *IO) Open(filename string) {
	io.status = IOStatusOpen
	io.filename = filename
	io.totalChars = 100
}

func (io *IO) Read() {
	if io.status != IOStatusOpen {
		io.status = IOStatusError
		io.readCallback("", errors.New("unable to open file"))
		return
	}

	if io.currentPosition >= io.totalChars {
		io.status = IOStatusClosed
		io.readCallback("", ErrEof)
		return
	}

	randomNumber := 5 + rand.Intn(5)
	chunk := fileContent[io.currentPosition : io.currentPosition+randomNumber+1]

	io.readCallback(chunk, nil)
	io.currentPosition += randomNumber + 1
}

func (io *IO) Close() {
	io.status = IOStatusClosed
}
