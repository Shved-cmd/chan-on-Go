package main

import "github.com/fatih/color"

type client chan<- string // an outgoing message channel

type clientInfo struct {
	ch   client
	name string
}

var (
	entering = make(chan clientInfo)
	leaving  = make(chan clientInfo)
	messages = make(chan string) // all incoming client messages
	green    = color.New(color.FgGreen).SprintfFunc()

	//
)

func main() {

}
