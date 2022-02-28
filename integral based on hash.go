package main

import (
	"fmt"
	"math"
)

func main() {

	f := func(x float64) float64 {
		return math.Pow(x, 2)
	}

	a := 0.0
	b := 2.0

	n := 100

	h := (b - a) / float64(n)

	integral := 0.0
	for i := 1; i <= n; i++ {
		integral += f(a + float64(i)*h)
	}

	fmt.Printf("The integral from %.2f to %.2f is %.2f\n", a, b, integral)
}
