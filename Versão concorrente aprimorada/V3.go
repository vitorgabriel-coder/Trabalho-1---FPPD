package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var forks [5]sync.Mutex 

func Philosopher(id, rounds int, wg *sync.WaitGroup) {
	time_to_think := 20 
	time_to_eat := 30   

	first := id
	second := (id + 1) % 5
	if first > second { 
		first, second = second, first
	}

	for iteration := 1; iteration <= rounds; iteration++ {
		
		fmt.Printf("%s T%d\n", strings.Repeat(" ", id*6), iteration) 
		time.Sleep(time.Duration(time_to_think) * time.Millisecond)

		
		forks[first].Lock()
		forks[second].Lock()

	
		fmt.Printf("%s E%d\n", strings.Repeat(" ", id*6), iteration) 
		time.Sleep(time.Duration(time_to_eat) * time.Millisecond)

	
		forks[second].Unlock()
		forks[first].Unlock()
	}
	wg.Done()
}

func main() {
	philosophers := 5
	rounds := 20
	var wg sync.WaitGroup

	fmt.Print("\n[P1] [P2] [P3] [P4] [P5]\n\n")

	
	start := time.Now()

	for j := 0; j < philosophers; j++ {
		wg.Add(1)
		go Philosopher(j, rounds, &wg)
	}
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("\nOrdered Dinner took %s\n\n", elapsed)
}