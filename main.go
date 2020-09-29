package main

import "fmt"

func main() {
	var n int
	var e float64
	e = 0
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		if i > 1 {
			nFactorial := factorial(i)
			e = e + (1 / float64(nFactorial))
		} else {
			e = e + (1 / 1)
		}
	}
	fmt.Println(e)

}

func factorial(n int) int {
	factorial := n
	for i := n; i > 1; i-- {
		factorial = factorial * (i - 1)
	}
	return factorial
}
