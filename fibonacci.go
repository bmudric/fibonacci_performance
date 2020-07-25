package main

func fibonacciRecursion(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return fibonacciRecursion(n - 1) + fibonacciRecursion(n - 2)
}