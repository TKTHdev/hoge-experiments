package main

func DotNaive(a, b []float32) float32 {
	sum := float32(0)
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func DotUnroll4(a, b []float32) float32 {
	sum := float32(0)
	for i := 0; i < len(a); i += 4 {
		s0 := a[i] * b[i]
		s1 := a[i+1] * b[i+1]
		s2 := a[i+2] * b[i+2]
		s3 := a[i+3] * b[i+3]
		sum += s0 + s1 + s2 + s3
	}
	return sum
}

func DotUnroll8(a, b []float32) float32 {
	sum := float32(0)
	for i := 0; i < len(a); i += 8 {
		s0 := a[i] * b[i]
		s1 := a[i+1] * b[i+1]
		s2 := a[i+2] * b[i+2]
		s3 := a[i+3] * b[i+3]
		s4 := a[i+4] * b[i+4]
		s5 := a[i+5] * b[i+5]
		s6 := a[i+6] * b[i+6]
		s7 := a[i+7] * b[i+7]
		sum += s0 + s1 + s2 + s3 + s4 + s5 + s6 + s7
	}
	return sum
}

func DotUnroll4Optimized(a, b []float32) float32 {
	sum0 := float32(0)
	sum1 := float32(0)
	sum2 := float32(0)
	sum3 := float32(0)

	length := len(a) & ^3

	for i := 0; i < length; i += 4 {
		_ = a[i+3]
		_ = b[i+3]
		sum0 += a[i] * b[i]
		sum1 += a[i+1] * b[i+1]
		sum2 += a[i+2] * b[i+2]
		sum3 += a[i+3] * b[i+3]
	}

	total := sum0 + sum1 + sum2 + sum3

	for i := length; i < len(a); i++ {
		total += a[i] * b[i]
	}
	return total
}

func Dot2Concurrent(a, b []float32) float32 {
	numGoroutines := 2
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := float32(0)
			for j := start; j < end; j++ {
				partialSum += a[j] * b[j]
			}
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func Dot4Concurrent(a, b []float32) float32 {
	numGoroutines := 4
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := float32(0)
			for j := start; j < end; j++ {
				partialSum += a[j] * b[j]
			}
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func Dot8Concurrent(a, b []float32) float32 {
	numGoroutines := 8
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := float32(0)
			for j := start; j < end; j++ {
				partialSum += a[j] * b[j]
			}
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func Dot16Concurrent(a, b []float32) float32 {
	numGoroutines := 16
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := float32(0)
			for j := start; j < end; j++ {
				partialSum += a[j] * b[j]
			}
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func DotAVX2(a, b []float32) float32

func Dot4ConcurrentAVX2(a, b []float32) float32 {
	numGoroutines := 4
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := DotAVX2(a[start:end], b[start:end])
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func Dot8ConcurrentAVX2(a, b []float32) float32 {
	numGoroutines := 8
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := DotAVX2(a[start:end], b[start:end])
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func Dot16ConcurrentAVX2(a, b []float32) float32 {
	numGoroutines := 16
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := DotAVX2(a[start:end], b[start:end])
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}

func Dot32ConcurrentAVX2(a, b []float32) float32 {
	numGoroutines := 32
	length := len(a)
	blockPerGoRoutine := length / numGoroutines
	results := make(chan float32, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * blockPerGoRoutine
		end := start + blockPerGoRoutine
		if i == numGoroutines-1 {
			end = length
		}
		go func(start, end int) {
			partialSum := DotAVX2(a[start:end], b[start:end])
			results <- partialSum
		}(start, end)
	}
	totalSum := float32(0)
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-results
	}
	return totalSum
}
