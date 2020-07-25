package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var fb = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
var fb10 uint64 = 55
var fb20 uint64 = 6765
var fb30 uint64 = 832040
var fb40 uint64 = 102334155
var fb50 uint64 = 12586269025
var fb60 uint64 = 1548008755920
var fb70 uint64 = 190392490709135
var fb80 uint64 = 23416728348467685
var fb90 uint64 = 2880067194370816120

func TestFibonacciRecursion(t *testing.T) {
	for n, expected := range fb {
		result := fibonacciRecursion(uint64(n))
		assert.Equal(t, expected, result)
	}
}

func TestFibonacciRecursion_10(t *testing.T) {
	result := fibonacciRecursion(10)
	assert.Equal(t, fb10, result)
}

func TestFibonacciRecursion_20(t *testing.T) {
	result := fibonacciRecursion(20)
	assert.Equal(t, fb20, result)
}

// PASS: TestFibonacciRecursion_30 (0.01s)
func TestFibonacciRecursion_30(t *testing.T) {
	result := fibonacciRecursion(30)
	assert.Equal(t, fb30, result)
}

// PASS: TestFibonacciRecursion_40 (0.78s)
func TestFibonacciRecursion_40(t *testing.T) {
	result := fibonacciRecursion(40)
	assert.Equal(t, fb40, result)
}

// PASS: TestFibonacciRecursion_50 (89.11s)
func TestFibonacciRecursion_50(t *testing.T) {
	result := fibonacciRecursion(50)
	assert.Equal(t, fb50, result)
}

// BenchmarkFibonacciRecursion_40
// BenchmarkFibonacciRecursion_40-8   	       2	 706926721 ns/op
func BenchmarkFibonacciRecursion_40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursion(40)
	}
}
