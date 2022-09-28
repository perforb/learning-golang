package main

import (
	"fmt"
	"testing"
)

// === RUN   TestIntMinBasic
// --- PASS: TestIntMinBasic (0.00s)
// === RUN   TestIntMinTableDriven
// === RUN   TestIntMinTableDriven/0,1
// === RUN   TestIntMinTableDriven/1,0
// === RUN   TestIntMinTableDriven/2,-2
// === RUN   TestIntMinTableDriven/0,-1
// === RUN   TestIntMinTableDriven/-1,0
// --- PASS: TestIntMinTableDriven (0.00s)
//     --- PASS: TestIntMinTableDriven/0,1 (0.00s)
//     --- PASS: TestIntMinTableDriven/1,0 (0.00s)
//     --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
//     --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
//     --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
// PASS
// ok      leaning-golang/gobyexample/testing-and-benchmarking     0.510s

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// goos: darwin
// goarch: arm64
// pkg: leaning-golang/gobyexample/testing-and-benchmarking
// BenchmarkIntMin-10      1000000000               0.3160 ns/op
// PASS
// ok      leaning-golang/gobyexample/testing-and-benchmarking     0.951s
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
