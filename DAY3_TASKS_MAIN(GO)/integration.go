package main

import (
	"fmt"
	"math"
	"runtime"

	"gonum.org/v1/gonum/integrate/quad"
	"gonum.org/v1/gonum/stat/distuv"
)

func Gaussain() {
	fmt.Println("Evaluate the expected value of x^2 + 3 under a Weibull distribution")
	f := func(x float64) float64 {
		d := distuv.Weibull{Lambda: 1, K: 1.5}
		return (x*x + 3) * d.Prob(x)
	}
	ev := quad.Fixed(f, 0, math.Inf(1), 10, nil, 0)
	fmt.Printf("EV with 10 points = %0.6v\n", ev)

	ev = quad.Fixed(f, 0, math.Inf(1), 30, nil, 0)
	fmt.Printf("EV with 30 points = %0.6v\n", ev)

	ev = quad.Fixed(f, 0, math.Inf(1), 100, nil, 0)
	fmt.Printf("EV with 100 points = %0.6v\n", ev)

	ev = quad.Fixed(f, 0, math.Inf(1), 10000, nil, 0)
	fmt.Printf("EV with 10000 points = %0.6v\n\n", ev)

	fmt.Println("Estimate using parallel evaluations of f.")
	concurrent := runtime.GOMAXPROCS(0)
	ev = quad.Fixed(f, 0, math.Inf(1), 100, nil, concurrent)
	fmt.Printf("EV = %0.6v\n", ev)
}
