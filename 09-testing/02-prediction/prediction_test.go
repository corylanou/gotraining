// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/VxnL7AEZSl

// go build -gcflags -m

// Package prediction provides code to show how branch prediction can affect performance.
package prediction

import (
	"math/rand"
	"testing"
)

func BenchmarkPredictable(b *testing.B) {
	data := make([]byte, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crunch(data)
	}
}

func BenchmarkUnpredictable(b *testing.B) {
	data := make([]byte, 1024)
	rand.Seed(0)

	// fill data with (pseudo) random noise
	for i := range data {
		// comment to make crunch behave predictably again
		data[i] = uint8(rand.Uint32())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crunch(data)
	}
}

func crunch(data []uint8) uint8 {
	var sum uint8
	for _, v := range data {
		if v < 128 {
			sum--
		} else {
			sum++
		}
	}
	return sum
}
