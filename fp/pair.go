package fp

type Pair[A, B any] interface {
	GetA() A
	GetB() B
}

type pair[A, B any] struct {
	a A
	b B
}

func PairOf[A, B any](a A, b B) Pair[A, B] {
	return pair[A, B]{a, b}
}

func (this pair[A, _]) GetA() A {
	return this.a
}

func (this pair[_, B]) GetB() B {
	return this.b
}
