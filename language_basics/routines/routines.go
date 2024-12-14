package main

import (
	"fmt"
	"time"
)

func main() {
	// making a channel
	endChannel := make(chan bool)
	go printInstantly("Hello", endChannel)
	go printInstantly("there", endChannel)
	go printWithDelay("wow, i was here with delay", endChannel)
	go printInstantly("hello there, im final", endChannel)
	<-endChannel
	<-endChannel
	<-endChannel
	<-endChannel
}

// / print instantly
func printInstantly(value string, endChannel chan bool) {
	fmt.Println(value)
	endChannel <- true
}

// print something with delay
func printWithDelay(value string, endChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Printed with delay: ", value)
	endChannel <- true
}
