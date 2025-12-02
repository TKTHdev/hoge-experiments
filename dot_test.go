package main

import "testing"

const N = 102400000

func BenchmarkDotNaive(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	DotNaive(x, y)
}

func BenchmarkDotUnroll4(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	DotUnroll4(x, y)
}

func BenchmarkDotUnroll8(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	DotUnroll8(x, y)
}

func BenchmarkDotUnroll4Optimized(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	DotUnroll4Optimized(x, y)
}

func BenchmarkDotAVX2(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	DotAVX2(x, y)
}

func BenchmarkDotConcurrent(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	DotConcurrent(x, y)
}
