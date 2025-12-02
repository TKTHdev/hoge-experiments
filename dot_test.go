package main

import "testing"

const N = 102400000

func BenchmarkDotNaive(b *testing.B){
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.StartTimer()
	DotNaive(x, y)
	b.ResetTimer()
}

func BenchmarkDotUnroll4(b *testing.B){
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.StartTimer()
	DotUnroll4(x, y)
	b.ResetTimer()
}

func BenchmarkDotUnroll8(b *testing.B){
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.StartTimer()
	DotUnroll8(x, y)
	b.ResetTimer()
}


func BenchmarkDotUnroll4Optimized(b *testing.B){
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.StartTimer()
	DotUnroll4Optimized(x, y)
	b.ResetTimer()
}