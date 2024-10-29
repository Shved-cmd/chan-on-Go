package main

type client chan<- string // an outgoing message channel

type clientInfo struct {
	ch   client
	name string
}

func main() {

}
