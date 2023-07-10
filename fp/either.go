package fp

type Either[L, R any] interface {
	IsLeft() bool
	IsRight() bool
	GetLeft() Option[L]
	GetRight() Option[R]
	Fold(func(L), func(R))
	Swap() Either[R, L]
}

type either[L, R any] struct {
	left  Option[L]
	right Option[R]
}

func LeftOf[L, R any](value L) Either[L, R] {
	return either[L, R]{
		left:  SomeOf(value),
		right: None[R](),
	}
}

func RightOf[L, R any](value R) Either[L, R] {
	return either[L, R]{
		left:  None[L](),
		right: SomeOf(value),
	}
}

func (x either[L, R]) IsLeft() bool {
	return x.left.IsDefined()
}

func (x either[L, R]) IsRight() bool {
	return x.right.IsDefined()
}

func (x either[L, R]) GetLeft() Option[L] {
	return x.left
}

func (x either[L, R]) GetRight() Option[R] {
	return x.right
}

func (x either[L, R]) Fold(left func(L), right func(R)) {
	if x.IsLeft() {
		v, _ := x.GetLeft().Get()
		left(v)
	} else {
		v, _ := x.GetRight().Get()
		right(v)
	}
}

func (x either[L, R]) Swap() Either[R, L] {
	return either[R, L]{
		left:  x.right,
		right: x.left,
	}
}

func MapEither[L, A, B any](x Either[L, A], f func(A) B) Either[L, B] {
	if x.IsRight() {
		v, _ := x.GetRight().Get()
		return RightOf[L, B](f(v))
	} else {
		v, _ := x.GetLeft().Get()
		return LeftOf[L, B](v)
	}
}

func FlatMapEither[L, A, B any](x Either[L, A], f func(A) Either[L, B]) Either[L, B] {
	if x.IsRight() {
		v, _ := x.GetRight().Get()
		return f(v)
	} else {
		v, _ := x.GetLeft().Get()
		return LeftOf[L, B](v)
	}
}
