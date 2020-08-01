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
//func TestFibonacciRecursion_50(t *testing.T) {
//	result := fibonacciRecursion(50)
//	assert.Equal(t, fb50, result)
//}

// BenchmarkFibonacciRecursion_40-8   	       2	 706926721 ns/op
func BenchmarkFibonacciRecursion_40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursion(40)
	}
}

func TestFibonacciLoop(t *testing.T) {
	for n, expected := range fb {
		result := fibonacciLoop(n)
		assert.Equal(t, expected, result)
	}
}

func TestFibonacciLoop_10(t *testing.T) {
	result := fibonacciLoop(10)
	assert.Equal(t, fb10, result)
}

func TestFibonacciLoop_20(t *testing.T) {
	result := fibonacciLoop(20)
	assert.Equal(t, fb20, result)
}

func TestFibonacciLoop_30(t *testing.T) {
	result := fibonacciLoop(30)
	assert.Equal(t, fb30, result)
}

// PASS: TestFibonacciLoop_40 (0.00s)
func TestFibonacciLoop_40(t *testing.T) {
	result := fibonacciLoop(40)
	assert.Equal(t, fb40, result)
}

// PASS: TestFibonacciLoop_50 (0.00s)
func TestFibonacciLoop_50(t *testing.T) {
	result := fibonacciLoop(50)
	assert.Equal(t, fb50, result)
}

// PASS: TestFibonacciLoop_60 (0.00s)
func TestFibonacciLoop_60(t *testing.T) {
	result := fibonacciLoop(60)
	assert.Equal(t, fb60, result)
}

// PASS: TestFibonacciLoop_70 (0.00s)
func TestFibonacciLoop_70(t *testing.T) {
	result := fibonacciLoop(70)
	assert.Equal(t, fb70, result)
}

// PASS: TestFibonacciLoop_80 (0.00s)
func TestFibonacciLoop_80(t *testing.T) {
	result := fibonacciLoop(80)
	assert.Equal(t, fb80, result)
}

// PASS: TestFibonacciLoop_90 (0.00s)
func TestFibonacciLoop_90(t *testing.T) {
	result := fibonacciLoop(90)
	assert.Equal(t, fb90, result)
}

// BenchmarkFibonacciLoop_40-8   	36849328	        28.1 ns/op
func BenchmarkFibonacciLoop_40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciLoop(40)
	}
}

// BenchmarkFibonacciLoop_90-8   	16818535	        60.9 ns/op
func BenchmarkFibonacciLoop_90(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciLoop(90)
	}
}

// PASS: TestFibonacciBinet (0.00s)
func TestFibonacciBinet(t *testing.T) {
	for n, expected := range fb {
		result := fibonacciBinet(n)
		assert.Equal(t, expected, result)
	}
}

// PASS: TestFibonacciBinet_50 (0.00s)
func TestFibonacciBinet_50(t *testing.T) {
	result := fibonacciBinet(50)
	assert.Equal(t, fb50, result)
}

//Error:      	Not equal:
//expected: 0x533163ef0321e5
//actual  : 0x533163ef0321dc
//Test:       	TestFibonacciBinet_80
//--- FAIL: TestFibonacciBinet_80 (0.00s)
func TestFibonacciBinet_80(t *testing.T) {
	result := fibonacciBinet(80)
	assert.Equal(t, fb80, result)
}

//Error:      	Not equal:
//expected: 0x27f80ddaa1ba7878
//actual  : 0x27f80ddaa1ba7600
//Test:       	TestFibonacciBinet_90
//--- FAIL: TestFibonacciBinet_90 (0.00s)
func TestFibonacciBinet_90(t *testing.T) {
	result := fibonacciBinet(90)
	assert.Equal(t, fb90, result)
}

// BenchmarkFibonacciBinet_40-8   	24540102	        42.4 ns/op
func BenchmarkFibonacciBinet_40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciBinet(40)
	}
}

// BenchmarkFibonacciBinet_90-8   	27428461	        43.5 ns/op
func BenchmarkFibonacciBinet_90(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciBinet(90)
	}
}