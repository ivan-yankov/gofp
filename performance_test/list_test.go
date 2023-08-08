package performance_test

import (
	"testing"
	"time"

	"github.com/ivan-yankov/gofp/fp"
)

func BenchmarkListForEachPar(b *testing.B) {
	n := 10
	seq := fp.ListTabulate(n, func(i int) int { return i + 1 })
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
