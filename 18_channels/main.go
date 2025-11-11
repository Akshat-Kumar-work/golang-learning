package main

import (
	"fmt"
	"math/rand"
)

//channels are use for communication between goroutines
//channels are like async queue or await until we get result

func processNum(numChan chan int) {
	//receiving number via numchan
	for num := range numChan {
		fmt.Println("processing number", num)
	}
	fmt.Println("processing number", <-numChan)
}

func main() {

	//creating channel to communicate b/w goroutines
	//in this case communication b/w main goroutine and processNum goroutine
	numChan := make(chan int)

	go processNum(numChan) //start another lightweight thread / goroutine

	//main goroutine/thread sends to numChan thread/goroutine
	numChan <- 5 //send value or message to that thread/goroutine

	for {
		numChan <- rand.Intn(100)
	}

	//.Sleep(time.Second * 2)

}
