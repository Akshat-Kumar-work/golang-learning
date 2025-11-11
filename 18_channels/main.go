package main

import (
	"fmt"
	"math/rand"
)

//NOTE : BELLOW ARE ALL UN-BUFFERED CHANNEL EXAMPLES
// WE CAN RECEIVE THE DATA BY CHANNEL ONE BY ONE SEQUENTIALY/SYNCRONISE WAY

// channels are use for communication between goroutines
// channels are like async queue or await until we get result
// NOTE: sending and receving via channel are blocking thing means it can block the thread
func processNum(numChan chan int) {
	//receiving number via numchan
	for num := range numChan {
		fmt.Println("processing number", num)
	}
	//receiving data from channel via <- channel name
	fmt.Println("processing number", <-numChan)
}

//we can also use channels to block the main thread
// as we already know the sending and receving is blocking
// so we can send any value or message
//once the goroutine is completed
// receive at main goroutine/thread
//so our main thread will wait until and unless it get some value via channel and it will get the values only when goroutine completed

// we can use sync.WaitGroup for multiple goroutines and channels for single goroutines
func task(done chan bool) {
	//using defer to schedule call this function just before when our task function return anything
	defer func() {
		done <- true
	}()
	fmt.Println("processing...")
}

func main() {

	//create done channel
	done := make(chan bool)
	//call task as goroutine
	go task(done)

	//receive the value send by done channel once the task goroutine completed
	<-done

	//creating channel to communicate b/w goroutines
	//in this case communication b/w main goroutine and processNum goroutine
	numChan := make(chan int)

	go processNum(numChan) //start another lightweight thread / goroutine

	//main goroutine/thread sends to numChan thread/goroutine
	numChan <- 5 //send value or message to that thread/goroutine via numChan (channel)

	for {
		numChan <- rand.Intn(100)
	}
	//close(numChan) //close the channel once you are done with sending data

	//.Sleep(time.Second * 2)

}
