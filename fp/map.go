package fp

import "reflect"

type IMap[K comparable, V any] interface {
	getData() map[K]V
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

type iMap[K comparable, V any] struct {
	data map[K]V
}

func mapCopy[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func EmptyMap[K comparable, V any]() IMap[K, V] {
	return iMap[K, V]{make(map[K]V)}
}

func MapOf[K comparable, V any](m map[K]V) IMap[K, V] {
	return iMap[K, V]{m}
}

func (this iMap[K, V]) getData() map[K]V {
	return mapCopy(this.data)
}

func (this iMap[K, V]) Put(key K, value V) IMap[K, V] {
	result := this.getData()
	result[key] = value
	return MapOf(result)
}

func (this iMap[K, V]) PutAll(that IMap[K, V]) IMap[K, V] {
	result := this.getData()
	for k, v := range that.getData() {
		result[k] = v
	}
	return MapOf(result)
}

func (this iMap[K, V]) Equals(that IMap[K, V]) bool {
	return reflect.DeepEqual(this, that)
}

func (this iMap[K, V]) Size() int {
	return len(this.data)
}

func (this iMap[K, V]) IsEmpty() bool {
	return this.Size() == 0
}

func (this iMap[K, V]) NonEmpty() bool {
	return !this.IsEmpty()
}

func (this iMap[K, V]) Get(key K) Option[V] {
	v, exists := this.data[key]
	if exists {
		return SomeOf(v)
	}
	return None[V]()
}

func (this iMap[K, V]) GetOrElse(key K, defaultValue V) V {
	v, exists := this.data[key]
	if exists {
		return v
	}
	return defaultValue
}

func (this iMap[K, V]) Remove(key K) IMap[K, V] {
	result := this.getData()
	delete(result, key)
	return MapOf(result)
}

func (this iMap[K, V]) ContainsKey(key K) bool {
	_, result := this.data[key]
	return result
}

func (this iMap[K, V]) ContainsValue(value V) bool {
	return this.Values().
		Exists(func(x V) bool { return reflect.DeepEqual(x, value) })
}

func (this iMap[K, V]) Entries() Seq[Pair[K, V]] {
	entries := []Pair[K, V]{}
	for key := range this.data {
		entries = append(entries, PairOf(key, this.data[key]))
	}
	return ArrayOfGoSlice(entries)
}

func (this iMap[K, V]) Keys() Seq[K] {
	return SeqMap(this.Entries(), func(x Pair[K, V]) K { return x.GetA() })
}

func (this iMap[K, V]) Values() Seq[V] {
	return SeqMap(this.Entries(), func(x Pair[K, V]) V { return x.GetB() })
}
