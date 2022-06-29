package main

import (
	"testing"
)

func Benchmark_parseLine(b *testing.B) {
	sample := make([][]string, 0)
	for i := 0; i <= 2500000; i++ {
		value := "000000,1111111,2222222"
		// logic to create first 3 columns
		temp := make([]string, 0)
		temp = append(temp, value)

		sample = append(sample, [][]string{temp}...)
	}

	for i := 0; i < b.N; i++ {
		parseLine(sample, 0)
	}
}
