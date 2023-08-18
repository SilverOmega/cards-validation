package cards_validation

import (
	_ "strconv"
	"testing"
	_ "testing"
	_ "time"
)

func BenchmarkSlienceNonAllocate(b *testing.B) {
	n := 100
	b.Run("test Benchmark Slice Non Allocate", func(b *testing.B) {
		//initializing
		for i := 0; i < b.N; i++ {
			ints := []int{}
			for i := 0; i < n; i++ {
				ints = append(ints, i)
			}
		}
	})

	b.Run("test Benchmark Slice Allocate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// benchmark code
			ints := make([]int, n)
			for i := 0; i < n; i++ {
				ints[i] = i
			}
		}
	})
}
