package main

import "fmt"

type Unary func(float64) float64
type UnaryMaker func(float64) Unary

// Functional refactorings:
//   Tennet principle
//   Binding
//   Wrap function
//   Inline definition

func main() {
	compose := func(f Unary, g Unary) Unary {
		return Unary(func(x float64) float64 {
			return f(g(x))
		})
	}

	make_adder := UnaryMaker(func(n float64) Unary {
		return Unary(func(x float64) float64 {
			return x + n
		})
	})

	make_multiplier := UnaryMaker(func(n float64) Unary {
		return Unary(func(x float64) float64 {
			return x * n
		})
	})

	fmt.Println(compose(make_multiplier(3), make_adder(1))(10))
}
