package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
)

func run(b *testing.B, f func()) {
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkListSize(b *testing.B) {
	run(b, func() { fp.ListFill(1000000, 1) })
}
