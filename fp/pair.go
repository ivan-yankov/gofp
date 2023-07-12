package fp

type Pair[A, B any] interface {
	GetA() A
	GetB() B
}

type pair[A, B any] struct {
	a A
	b B
}

func (this pair[A, _]) GetA() A {
	return this.a
}

func (this pair[_, B]) GetB() B {
	return this.b
}
