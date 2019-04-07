package main

import "fmt"

type Unary func(float64) float64

func main() {
	compose := func(f Unary, g Unary) Unary {
		return Unary(func(x float64) float64 {
			return f(g(x))
		})
	}

	make_adder := func(n float64) Unary {
		return Unary(func(x float64) float64 {
			return x + n
		})
	}

	make_multiplier := func(n float64) Unary {
		return Unary(func(x float64) float64 {
			return x * n
		})
	}

	fmt.Println(compose(make_multiplier(3), make_adder(1))(10))
}
