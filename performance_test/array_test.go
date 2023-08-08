package performance_test

import (
	"testing"
	"time"

	"github.com/ivan-yankov/gofp/fp"
)

func BenchmarkArrayFindSlice(b *testing.B) {
	na := 10000
	nb := 100
	la := fp.ArrayTabulate(na, func(i int) int { return i + 1 })
	lb := fp.ArrayTabulate(nb, func(i int) int { return i + na - nb - 10 })
	b.ResetTimer()
	run(b, func() { la.FindSlice(lb) })
}

func BenchmarkArrayForEachPar(b *testing.B) {
	n := 10
	seq := fp.ArrayTabulate(n, func(i int) int { return i + 1 })
	b.ResetTimer()
	run(b, func() {
		seq.ForEachPar(
			func(int) fp.Unit {
				time.Sleep(time.Second)
				return fp.GetUnit()
			},
		)
	})
}
