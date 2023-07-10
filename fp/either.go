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

func (this either[L, R]) IsLeft() bool {
	return this.left.IsDefined()
}

func (this either[L, R]) IsRight() bool {
	return this.right.IsDefined()
}

func (this either[L, R]) GetLeft() Option[L] {
	return this.left
}

func (this either[L, R]) GetRight() Option[R] {
	return this.right
}

func (this either[L, R]) Fold(left func(L), right func(R)) {
	if this.IsLeft() {
		v, _ := this.GetLeft().Get()
		left(v)
	} else {
		v, _ := this.GetRight().Get()
		right(v)
	}
}

func (this either[L, R]) Swap() Either[R, L] {
	return either[R, L]{
		left:  this.right,
		right: this.left,
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
