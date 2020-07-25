package main

func fibonacciRecursion(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return fibonacciRecursion(n - 1) + fibonacciRecursion(n - 2)
}

func fibonacciLoop(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}
	var x = uint64(0)
	var y = uint64(1)
	var sum = uint64(0)
	for i := 0; i < n - 1; i++ {
		sum = x + y
		x = y
		y = sum
	}
	return sum
}