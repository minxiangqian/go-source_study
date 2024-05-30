package atomic

import (
	"testing"
)

func TestAtomicAddVal(t *testing.T) {
	atomicAddVal()
}
func TestMutexAddVal(t *testing.T) {
	mutexAddVal()
}

func TestAddVal(t *testing.T) {
	addVal()
}

func BenchmarkAtomicVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicAddVal()
	}
}
func BenchmarkMutexVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutexAddVal()
	}
}

func BenchmarkAddVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addVal()
	}
}
