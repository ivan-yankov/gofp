package fp

type List[T any] struct {
	head  T
	tail  Seq[T]
	empty bool
}

func emptyList[T any]() Seq[T] {
	return List[T]{
		head:  *new(T),
		tail:  nil,
		empty: true,
	}
}

func ListOf[T any](elements ...T) Seq[T] {
	var loop func(index int, acc Seq[T]) Seq[T]
	loop = func(index int, acc Seq[T]) Seq[T] {
		if index < 0 {
			return acc
		} else {
			return loop(index-1, acc.Add(elements[index]))
		}
	}

	if len(elements) != 0 {
		return loop(len(elements)-1, emptyList[T]())
	} else {
		return emptyList[T]()
	}
}

func (x List[T]) IsEmpty() bool {
	return x.empty
}

func (x List[T]) NonEmpty() bool {
	return !x.empty
}

func (x List[T]) HeadOption() Option[T] {
	if x.NonEmpty() {
		return SomeOf(x.head)
	} else {
		return None[T]()
	}
}

func (x List[T]) LastOption() Option[T] {
	var iterate func(Seq[T]) T
	iterate = func(seq Seq[T]) T {
		if seq.Tail().IsEmpty() {
			v, _ := seq.HeadOption().Get()
			return v
		} else {
			return iterate(seq.Tail())
		}
	}

	if x.NonEmpty() {
		return SomeOf(iterate(x))
	} else {
		return None[T]()
	}
}

func (x List[T]) Tail() Seq[T] {
	if x.tail == nil {
		return emptyList[T]()
	} else {
		return x.tail
	}
}

func (x List[T]) Add(e T) Seq[T] {
	if x.NonEmpty() {
		return List[T]{
			head: e,
			tail: List[T]{
				head:  x.head,
				tail:  x.tail,
				empty: false,
			},
			empty: false,
		}
	}

	return List[T]{
		head:  e,
		tail:  nil,
		empty: false,
	}
}
