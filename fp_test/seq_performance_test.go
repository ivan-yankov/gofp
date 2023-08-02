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

func BenchmarkListFindSlice(b *testing.B) {
	na := 100000
	nb := 50
	la := fp.ListTabulate(na, func(i int) int { return i + 1 })
	lb := fp.ListTabulate(nb, func(i int) int { return i + na - nb - 10 })
	b.ResetTimer()
	run(b, func() { fp.SeqFindSlice(la, lb) })
}

func BenchmarkArrayFindSlice(b *testing.B) {
	na := 100000
	nb := 50
	la := fp.ArrayTabulate(na, func(i int) int { return i + 1 })
	lb := fp.ArrayTabulate(nb, func(i int) int { return i + na - nb - 10 })
	b.ResetTimer()
	run(b, func() { fp.SeqFindSlice(la, lb) })
}
