package counter

import (
	"testing"
)

func BenchmarkBasic_Inc(b *testing.B) {
	counter := NewBasicCounter()
	benchInc(b, counter)
}

func BenchmarkBasic_Count(b *testing.B) {
	counter := NewBasicCounter()
	benchCount(b, counter)
}

func BenchmarkMutex_Inc(b *testing.B) {
	counter := NewMutexCounter()
	benchInc(b, counter)
}

func BenchmarkMutex_Count(b *testing.B) {
	counter := NewMutexCounter()
	benchCount(b, counter)
}

func BenchmarkRWMutex_Inc(b *testing.B) {
	counter := NewRWMutexCounter()
	benchInc(b, counter)
}

func BenchmarkRWMutex_Count(b *testing.B) {
	counter := NewRWMutexCounter()
	benchCount(b, counter)
}

func BenchmarkAtomic_Inc(b *testing.B) {
	counter := NewAtomicCounter()
	benchInc(b, counter)
}

func BenchmarkAtomic_Count(b *testing.B) {
	counter := NewAtomicCounter()
	benchCount(b, counter)
}

func benchInc(b *testing.B, counter Counter) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.Inc()
	}
}

func benchCount(b *testing.B, counter Counter) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.Count()
	}
}
