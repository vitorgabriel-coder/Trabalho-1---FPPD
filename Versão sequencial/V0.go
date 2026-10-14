package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var waitTotal time.Duration
var waitMax time.Duration
var eatingNow int
var eatingMax int
var cycles [5]int

func addWait(d time.Duration) {
	mu.Lock()
	waitTotal += d
	if d > waitMax {
		waitMax = d
	}
	mu.Unlock()
}

func startEating() {
	mu.Lock()
	eatingNow++
	if eatingNow > eatingMax {
		eatingMax = eatingNow
	}
	mu.Unlock()
}

func endEating(id int) {
	mu.Lock()
	eatingNow--
	cycles[id]++
	mu.Unlock()
}

func printMetrics(elapsed time.Duration, philosophers, rounds int) {
	fmt.Printf("\nTempo total: %s\n", elapsed)
	fmt.Printf("Tempo médio de espera pelos garfos: %s\n", waitTotal/time.Duration(philosophers*rounds))
	fmt.Printf("Maior tempo de espera: %s\n", waitMax)
	fmt.Printf("Máximo de filósofos comendo ao mesmo tempo: %d\n", eatingMax)
	fmt.Printf("Ciclos concluídos por filósofo: %v\n\n", cycles)
}

func spaces(id int) string {
	s := ""
	for i := 1; i <= id*6; i++ {
		s += " "
	}
	return s
}

func Philosopher(id, iteration int) {
	time_to_think := 20
	time_to_eat := 30

	fmt.Printf("%s T%d\n", spaces(id), iteration)
	time.Sleep(time.Duration(time_to_think) * time.Millisecond)

	addWait(0)
	startEating()
	fmt.Printf("%s E%d\n", spaces(id), iteration)
	time.Sleep(time.Duration(time_to_eat) * time.Millisecond)
	endEating(id)
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

	fmt.Print("\nJantar Sequencial")
	printMetrics(elapsed, philosophers, rounds)
}