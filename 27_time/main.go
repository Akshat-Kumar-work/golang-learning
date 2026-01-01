package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	currTime := time.Now()

	fmt.Println("current Time: ", currTime)
	fmt.Println("current hour:", currTime.Hour())

	time.Sleep(3 * time.Second)

	//fires one single time by sending currTime over -> timer.c
	singleshotTimer := time.NewTimer(4 * time.Second)
	//fmt.Println("single shot timer", singleshotTimer)

	//fire repeatedly after provided time over -> timer.c
	periodicTimer := time.NewTicker(2 * time.Second)
	//fmt.Println("preiodic Timer", periodicTimer)

	//done := make(chan bool, 2)

	var wg sync.WaitGroup

	wg.Add(1)
	//goroutine for single shot timer
	go func() {
		defer wg.Done()
		t := <-singleshotTimer.C
		fmt.Println("single shot timer fired at", t)
		//done <- true
	}()

	wg.Add(1)
	//goroutine for periodic timer
	go func() {
		defer wg.Done()
		t := <-periodicTimer.C
		fmt.Println("preiodic timer fired at", t)
	}()

	//new syntax for wait group
	// wg.Go(func() {
	// 	t := <-singleshotTimer.C
	// 	fmt.Println("single shot timer fired at", t)
	// })

	wg.Go(func() {
		t := <-periodicTimer.C
		fmt.Println("periodic timer fired at", t)
	})

	wg.Wait()

	//<-done
	periodicTimer.Stop()
	fmt.Printf("timers completed")

}
