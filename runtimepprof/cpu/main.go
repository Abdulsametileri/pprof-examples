package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	expensiveFunc()
}

//go:noinline
func expensiveFunc() {
	label := pprof.Labels("expensiveFunc", "sum of values at length of 10m")

	pprof.Do(context.Background(), label, func(_ context.Context) {
		var sum float64
		for i := 0; i < 10_000_000; i++ {
			sum += rand.Float64()
		}
	})

	anotherExpensiveFunc()
}

//go:noinline
func anotherExpensiveFunc() {
	label := pprof.Labels("anotherExpensiveFunc", "sum of values at length of 1m")

	pprof.Do(context.Background(), label, func(ctx context.Context) {
		var sum int
		for i := 0; i < 1_000_000; i++ {
			sum += rand.Intn(10)
		}
	})
}
