package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	//defer keyword used to schedule a function call to be executed just before the surrounding function return
	defer wg.Done() // mark this task as done when finished or decrement the goroutine count
	fmt.Println("doing task", id)
}

// BY using go keyword we can get multithreading over any task/method/function
// go has its own scheduler it does not depend over os scheduler to schedule task
func main() {

	//sync.WaitGroup is a synchronization primitive in Go’s sync package
	//used to wait for multiple goroutines to finish before continuing execution
	// we use sync.WaitGroup only when we have multiple goroutine
	// we can say it is like promise.all in js or asyncio.gather in python to execute multiple async operation concurrently/parallely
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //goroutine counter for keeping track how many goroutines are running
		//we can say go is used to start concurrent task or parallel task if multicores
		// it works as async same as in python and js

		go task(i, &wg)
	}

	//using this wg so that our main thread does not get exit before task threads are working
	//as the task function run over different threads(cores)
	wg.Wait() // waits until all goroutines call wg.Done()
	fmt.Println("All tasks done!")
}

// 🟢 Go’s Goroutines

// Run within a single process.

// Share heap memory (no copying data between processes).

// Scheduler automatically assigns them to threads.

// Extremely lightweight (~2 KB stack vs ~MB for an OS thread).

// Switching between goroutines happens in user space (Go runtime), not OS-level context switching.

// ✅ Super fast context switching
// ✅ Can spawn millions of goroutines
// ✅ True parallel execution on multiple cores
// ⚠️ Must manage shared state carefully (race conditions)

// ✅ 1. Single Process

// When you run a Go program (go run main.go), it launches one OS process.

// Inside that single process, Go’s runtime manages everything: goroutines, threads, scheduling, garbage collection, etc.

// So there’s one Go process, but it can fully utilize all CPU cores.

// ✅ 2. Multiple OS Threads (Managed by Go Runtime)

// The Go runtime creates multiple OS threads (M in the GMP model).

// Your goroutines (G) — which are lightweight user-space “threads” — get scheduled across those OS threads.

// So you might have:
// 1 process
// ├── 100 OS threads
// │    ├── hundreds or thousands of goroutines running on them

// ✅ 3. True Parallelism

// Because Go doesn’t have a Global Interpreter Lock (GIL) (unlike Python), multiple OS threads can run Go code simultaneously on multiple CPU cores.

// The Go scheduler automatically distributes runnable goroutines across available cores — no manual thread management required.

// So yes — Go code can literally run in parallel across cores, inside one process.

// ✅ 4. No GIL or “Main Thread” Restriction

// There’s no GIL — Go’s runtime and garbage collector are designed to be thread-safe.

// There’s no single “main thread” bottleneck.
// The main() goroutine starts execution, but after that, any goroutine can run on any thread or core.

// Example:
// func main() {
//     go task1()
//     go task2()
//     fmt.Println("main finished")
// }

// All three (main, task1, task2) can run truly in parallel.

// ✅ 5. Goroutines ≠ OS Threads

// Goroutines are much lighter than threads (≈ 2 KB stack).

// Go’s M:N scheduler maps many goroutines (G) to fewer OS threads (M) dynamically.

// Context switching happens in user space, not via the OS → it’s fast.

//	NOTE : go has multithreading & multiprocessing , but it does not need multiprocessing for
// true parallelism like js and python because it does not have gil or 1 main thread
// it can spawn multiple thread in a single go process which can utilize multi cores of cpu
