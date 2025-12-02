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
	for i := 0; i < b.N; i++ {
		DotNaive(x, y)
	}
}

func BenchmarkDotUnroll4(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DotUnroll4(x, y)
	}
}

func BenchmarkDotUnroll8(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DotUnroll8(x, y)
	}
}

func BenchmarkDotUnroll4Optimized(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DotUnroll4Optimized(x, y)
	}
}

func BenchmarkDotAVX2(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DotAVX2(x, y)
	}
}

func BenchmarkDotVNNI(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DotVNNI(x, y)
	}
}

func BenchmarkDot2Concurrent(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot2Concurrent(x, y)
	}
}

func BenchmarkDot4Concurrent(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot4Concurrent(x, y)
	}
}

func BenchmarkDot8Concurrent(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot8Concurrent(x, y)
	}
}

func BenchmarkDot16Concurrent(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot16Concurrent(x, y)
	}
}

func BenchmarkDot4ConcurrentAVX2(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot4ConcurrentAVX2(x, y)
	}
}

func BenchmarkDot8ConcurrentAVX2(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot8ConcurrentAVX2(x, y)
	}
}

func BenchmarkDot16ConcurrentAVX2(b *testing.B) {
	x := make([]float32, N)
	y := make([]float32, N)
	for i := 0; i < N; i++ {
		x[i] = float32(i)
		y[i] = float32(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dot16ConcurrentAVX2(x, y)
	}
}