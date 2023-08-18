package cards_validation

import (
	"bytes"
	_ "strconv"
	"testing"
	_ "testing"
	_ "time"
)

func LastFour(msg []byte) {
	buf := new(bytes.Buffer)
	buf.Write(msg)
}

func BenchmarkCard_LastFour(b *testing.B) {
	msg := []byte("number is not long enough")
	for i := 0; i < b.N; i++ { // internal bench stuff handle
		for i := 0; i < 4; i++ {
			LastFour([]byte(msg))
		}
	}
}
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
