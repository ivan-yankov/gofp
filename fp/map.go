package fp

import "reflect"

type IMap[K, V any] interface {
	Put(K, V) IMap[K, V]
	PutAll(IMap[K, V]) IMap[K, V]
	Equals(IMap[K, V]) bool
	Size() int
	IsEmpty() bool
	NonEmpty() bool
	Get(K) Option[V]
	GetOrElse(K, V) V
	Remove(K) IMap[K, V]
	ContainsKey(K) bool
	ContainsValue(V) bool
	Entries() Seq[Pair[K, V]]
	Keys() Seq[K]
	Values() Seq[V]
}

type iMap[K, V any] struct {
	data Seq[Pair[K, V]]
}

func MapOf[K, V any](items ...Pair[K, V]) IMap[K, V] {
	return iMap[K, V]{ArrayOfGoSlice(items)}
}

func MapOfGoSlice[K, V any](items []Pair[K, V]) IMap[K, V] {
	return iMap[K, V]{ArrayOfGoSlice(items)}
}

func MapOfSeq[K, V any](seq Seq[Pair[K, V]]) IMap[K, V] {
	return iMap[K, V]{seq}
}

func (this iMap[K, V]) Put(key K, value V) IMap[K, V] {
	return MapOfSeq(this.Remove(key).Entries().Add(PairOf(key, value)))
}

func (this iMap[K, V]) PutAll(that IMap[K, V]) IMap[K, V] {
	thatKeys := that.Keys()
	return MapOfSeq(
		this.Entries().
			FilterNot(func(x Pair[K, V]) bool { return thatKeys.ContainsElement(x.GetA()) }).
			Concat(that.Entries()),
	)
}

func (this iMap[K, V]) Equals(that IMap[K, V]) bool {
	return SeqSetEquals(this.Entries(), that.Entries())
}

func (this iMap[K, V]) Size() int {
	return this.data.Size()
}

func (this iMap[K, V]) IsEmpty() bool {
	return this.Size() == 0
}

func (this iMap[K, V]) NonEmpty() bool {
	return !this.IsEmpty()
}

func (this iMap[K, V]) Get(key K) Option[V] {
	return OptionMap(
		this.Entries().
			Find(func(x Pair[K, V]) bool { return reflect.DeepEqual(x.GetA(), key) }),
		func(x Pair[K, V]) V { return x.GetB() },
	)
}

func (this iMap[K, V]) GetOrElse(key K, defaultValue V) V {
	return this.Get(key).GetOrElse(defaultValue)
}

func (this iMap[K, V]) Remove(key K) IMap[K, V] {
	return MapOfSeq(
		this.Entries().
			FilterNot(func(x Pair[K, V]) bool { return reflect.DeepEqual(x.GetA(), key) }),
	)
}

func (this iMap[K, V]) ContainsKey(key K) bool {
	return this.Entries().
		Exists(func(x Pair[K, V]) bool { return reflect.DeepEqual(x.GetA(), key) })
}

func (this iMap[K, V]) ContainsValue(value V) bool {
	return this.Entries().
		Exists(func(x Pair[K, V]) bool { return reflect.DeepEqual(x.GetB(), value) })
}

func (this iMap[K, V]) Entries() Seq[Pair[K, V]] {
	return this.data
}

func (this iMap[K, V]) Keys() Seq[K] {
	return SeqMap(this.Entries(), func(x Pair[K, V]) K { return x.GetA() })
}

func (this iMap[K, V]) Values() Seq[V] {
	return SeqMap(this.Entries(), func(x Pair[K, V]) V { return x.GetB() })
}
