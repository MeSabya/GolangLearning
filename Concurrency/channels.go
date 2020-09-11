package main

import "fmt"

//pings : is to send msg to channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//Here pings: receiving from a channel, pongs is receiving channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Message Passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
