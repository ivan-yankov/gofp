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
	list := fp.ListFill(1000000, 1)
	b.ResetTimer()
	run(b, func() { list.Size() })
}

func BenchmarkListContainsSlice(b *testing.B) {
	na := 100000
	nb := 50
	la := fp.ListTabulate(na, func(i int) int { return i + 1 })
	lb := fp.ListTabulate(nb, func(i int) int { return i + na - nb - 10 })
	b.ResetTimer()
	run(b, func() { la.FindSlice(lb) })
}
