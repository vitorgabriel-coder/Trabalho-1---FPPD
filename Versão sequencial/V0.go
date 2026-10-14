package main

import (
	"fmt"
	"time"
)

func Philosopher(id, iteration int) {
	time_to_think := 20 
	time_to_eat := 30   

	
	   for s := 1; s <= id*6; s++ {
          fmt.Printf(" ")
          }
     fmt.Printf(" T%d\n", iteration) 
     time.Sleep(time.Duration(time_to_think) * time.Millisecond)


	for s := 1; s <= id*6; s++ {
          fmt.Printf(" ")
          }
     fmt.Printf(" E%d\n", iteration) 
     time.Sleep(time.Duration(time_to_eat) * time.Millisecond)
}

func main() {
	philosophers := 5
	rounds := 20

	fmt.Print("\n[P1] [P2] [P3] [P4] [P5]\n\n")


	start := time.Now()

	for i := 1; i <= rounds; i++ {
		for j := 0; j < philosophers; j++ {
			Philosopher(j, i)
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("\nSequential Dinner took %s\n\n", elapsed)
}