package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	runtime.GC() // invoke gc, in order to get up-to-date statistics

	expensiveFunc()

	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}

//go:noinline
func expensiveFunc() {
	m := make([]int, 10_000_000)
	for i := range m {
		m[i] = rand.Intn(127)
	}

	anotherExpensiveFunc()
}

//go:noinline
func anotherExpensiveFunc() {
	m := make(map[int]float64, 1_000_000)

	for key := range m {
		m[key] = rand.Float64()
	}
}
