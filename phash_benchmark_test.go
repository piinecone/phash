package phash

import (
	"testing"
)

func BenchmarkHammingDistanceC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingDistanceC(10631249863007147461, 17752165113623953620)
	}
}

func BenchmarkHammingDistanceGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HammingDistance(10631249863007147461, 17752165113623953620)
	}
}
