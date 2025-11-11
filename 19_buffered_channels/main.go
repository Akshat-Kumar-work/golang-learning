package main

import (
	"fmt"
	"time"
)

// A buffered channel has space to hold a fixed number of values.

// A send only blocks when the buffer is full.
// A receive only blocks when the buffer is empty.

// to mark channel as receive only or send only
// here emailChannel is receive only
// here done channel is send only
func emailSender(emailChannel <-chan string, done chan<- bool) {
	defer func() {
		done <- true
	}()
	//emailChannel <- "testing if we can send to emailChannel or not "
	for email := range emailChannel {
		fmt.Println("sending email to ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int, 2) // buffer of size 2

	ch <- 1 // doesn't block
	ch <- 2 // doesn't block
	//ch <- 3 // would block because buffer is full

	fmt.Println(<-ch) // reads one, frees space
	ch <- 3           // now this works

	close(ch) //always close a channel when you are done with sending values

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 1
		ch2 <- "ok"
	}()

	//receving values from multiple channels
	for i := 0; i < 2; i++ {
		select {
		case cha1Value := <-ch1:
			fmt.Println("channel 1 value", cha1Value)
		case cha2Value := <-ch2:
			fmt.Println("channel 2 value", cha2Value)
		}

	}

}
