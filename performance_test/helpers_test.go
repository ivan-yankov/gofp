package performance_test

import "testing"

func run(b *testing.B, f func()) {
	for i := 0; i < b.N; i++ {
		f()
	}
}
