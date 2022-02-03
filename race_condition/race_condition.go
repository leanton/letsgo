/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/

package main

import (
	"fmt"
	"sync"
)

/*
The application attached shows the race condition if the variable x is updated
and read at the same time. In the 1st goroutine we are observing the current
value of x by simply printing it. In the 2nd goroutine we are incrementing the
value of x 1_000_000 times

If you launch the program few times you will see that the result is
different
*/
func main() {
	var wg sync.WaitGroup

	var x int
	wg.Add(2)
	go printX(&wg, &x)
	go updateX(&wg, &x)
	wg.Wait()

}

func printX(wg *sync.WaitGroup, x *int) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(*x)
	}
}

func updateX(wg *sync.WaitGroup, x *int) {
	defer wg.Done()

	for i := 0; i < 1_000_000; i++ {
		*x = *x + 1
	}
}
