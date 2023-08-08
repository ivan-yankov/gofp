package performance_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
)

func BenchmarkListFindSlice(b *testing.B) {
	na := 10000
	nb := 100
	la := fp.ListTabulate(na, func(i int) int { return i + 1 })
	lb := fp.ListTabulate(nb, func(i int) int { return i + na - nb - 10 })
	b.ResetTimer()
	run(b, func() { fp.SeqFindSlice(la, lb) })
}

func BenchmarkArrayFindSlice(b *testing.B) {
	na := 10000
	nb := 100
	la := fp.ArrayTabulate(na, func(i int) int { return i + 1 })
	lb := fp.ArrayTabulate(nb, func(i int) int { return i + na - nb - 10 })
	b.ResetTimer()
	run(b, func() { fp.SeqFindSlice(la, lb) })
}
