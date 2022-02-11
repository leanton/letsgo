/*
Implement the dining philosopher’s problem with the following
constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between
each adjacent pair of philosophers.

2. Each philosopher should eat only 3 times (not in an infinite loop as we did
in lecture)

3. The philosophers pick up the chopsticks in any order, not lowest-numbered
first (which we did in lecture).

4. In order to eat, a philosopher must get
permission from a host which executes in its own goroutine.

5. The host allows no more than 2 philosophers to eat concurrently.

6. Each philosopher is numbered, 1 through 5.

7. When a philosopher starts eating (after it has obtained necessary locks) it
prints “starting to eat <number>” on a line by itself, where <number> is the
number of the philosopher.

8. When a philosopher finishes eating (before it has released its locks) it
prints “finishing eating <number>” on a line by itself, where <number> is the
number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 2)

	sticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, sticks[i], sticks[(i+1)%5]}
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, ch)
	}
	go coordinator(ch)
	wg.Wait()
}

func coordinator(ch chan int) {
	for range ch {
	}
}

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- p.id
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", p.id)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("finishing eating", p.id)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}
