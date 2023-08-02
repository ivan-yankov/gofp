package fp

import (
	"fmt"
)

func loop[T any](
	i int,
	inc func(int) int,
	p func(int) bool,
	f func(int, T) T,
	acc T) T {

	if p(i) {
		return acc
	}

	return loop(inc(i), inc, p, f, f(i, acc))
}

func emptyList[T any]() Seq[T] {
	return List[T]{
		head: *new(T),
		tail: nil,
		size: 0,
	}
}

func emptyArray[T any]() Seq[T] {
	return Array[T]{[]T{}}
}

func emptySeq[T any](list bool) Seq[T] {
	if list {
		return emptyList[T]()
	}
	return emptyArray[T]()
}

func findIndex[T any](seq Seq[T], p func(int, T, Option[int]) bool) int {
	f := func(i int, e T, acc Option[int]) Option[int] {
		if p(i, e, acc) {
			return SomeOf(i)
		}
		return acc
	}

	return SeqFoldCount[T, Option[int]](seq, f, None[int]()).GetOrElse(-1)
}

func add[T any](e T, acc Seq[T]) Seq[T] {
	return acc.Add(e)
}

func collect[T any](
	seq Seq[T],
	appendCondition func(int, T, Seq[T]) bool,
	emptySeq func() Seq[T]) Seq[T] {

	f := func(i int, e T, acc Seq[T]) Seq[T] {
		if appendCondition(i, e, acc) {
			return acc.Add(e)
		}
		return acc
	}

	return SeqFoldCount[T, Seq[T]](seq, f, emptySeq()).Reverse()
}

func minInt(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func maxInt(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func mkString[T any](seq Seq[T], sep string) string {
	strings := SeqMap[T, string](
		seq,
		func(x T) string { return fmt.Sprintf("%+v", x) },
	)
	lastIndex := seq.Size() - 1
	f := func(i int, e string, acc string) string {
		if i == lastIndex {
			return acc + e
		}
		return acc + e + sep
	}

	return SeqFoldCount[string, string](strings, f, "")
}

func prefixLength[T any](seq Seq[T], p func(T) bool) int {
	type Acc struct {
		n     int
		check bool
	}

	f := func(e T, acc Acc) Acc {
		if acc.check && p(e) {
			return Acc{acc.n + 1, true}
		}
		return Acc{acc.n, false}
	}

	r := SeqFoldLeft[T, Acc](seq, f, Acc{0, true})
	return r.n
}

func reduce[T any](seq Seq[T], f func(T, T) T) Option[T] {
	if seq.IsEmpty() {
		return None[T]()
	}

	r := SeqFoldLeft[T, T](seq.Tail(), f, seq.HeadOption().Get())
	return SomeOf(r)
}
