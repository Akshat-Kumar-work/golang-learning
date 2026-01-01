package main

import "context"

//context -> It is a request lifetime or lifecycle
// it is used for following
// cancellation
// timeouts
// deadlines
// request lifetime
// passing request-scoped data

// .........(1) contex.Context
// -> an interface that carries cancellation signal, deadline/timeout, request-scoped values
// context.Context contains -> Done(), Err(), Deadline() , Value()
// always pass context.Context as first paramter of function
// example -> func DoWork(ctx context.Context, id int) error
//The http.Request has a Context() method that returns a context.Context associated with the request. This context is canceled when the request is completed.

//.......... (2) context.Background()
// -> it is a function that create a root context
// never canceled, no timeout, no values
// used in main(), app startup, top-level workers
// example -> ctx := context.Background()

//......... (3) context.WithCancel() ->(manual stop)
// it is a function create a child context that can manually cancel
// ctx, cancel := context.WithCancel(parent)
//example
// ctx, cancel := context.WithCancel(context.Background())

// go func() {
//     time.Sleep(2 * time.Second)
//     cancel()
// }()

// select {
// case <-ctx.Done(): //will receive done when the context get cancell
//     fmt.Println("stopped")
// }

// ..........(4) context.WithTimeout() >(auto stop)
// it is function which automatically cancels after a duration
// example
// ctx, cancel := context.WithTimeout(parent, 3*time.Second)	stop after 3 seconds
// defer cancel()

//............(5) context.WithDeadline() -> cancel/stop at specific time
// example
//ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
//ctx, cancel := context.WithDeadline(parent, time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC))

// ...........(6) ctx.Done() ->(stop signal)
// A channel that closes when a context is cancelled
// every long-runnig goroutine must listen to ctx.Done()
// example
// for {
//     select {
//     case <-ctx.Done():
//         return
//     default:
//         doWork()
//     }
// }

func main() {
	ctx := context.Background()
	//context err handling
	select {
	case <-ctx.Done():
		// Check what caused cancellation
		switch ctx.Err() {
		case context.Canceled:
			// Manual cancellation
		case context.DeadlineExceeded:
			// Timeout
		}
	default:
		// Continue work
	}

}

//Imp

// Never store structs, maps, slices in context.Value() (immutable data only)

// Always call cancel() function (even with timeout) to release resources

// Never pass nil context - use context.Background() or context.TODO()

// Don't use context for passing optional function parameters
