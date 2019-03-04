package main

import (
	"log"
	"time"
	"math"
	"flag"
)
func main() {
	init_start := time.Now()
	iterations := flag.Int("n", 5000, "number of iterations")
	flag.Parse();
	log.Printf("Calculating pi with %d iterations", *iterations)

	start := time.Now()
	log.Printf("Value of pi = %2.30f", pi(*iterations))
	log.Printf("Time taken for pi calculation - %s", time.Since(start))
	log.Printf("Total time taken - %s", time.Since(init_start))
}

// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

