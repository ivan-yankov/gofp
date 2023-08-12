package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestMapPut(t *testing.T) {
	assert.Equal(t, fp.MapOf(map[int]string{1: "s1"}), fp.EmptyMap[int, string]().Put(1, "s1"))
	assert.Equal(t, fp.MapOf(map[int]string{1: "s2"}), fp.MapOf(map[int]string{1: "s1"}).Put(1, "s2"))
	assert.Equal(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}), fp.MapOf(map[int]string{1: "s1"}).Put(2, "s2"))
}

func TestMapPutAll(t *testing.T) {
	m := fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"})
	assert.Equal(t, fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"}), fp.EmptyMap[int, string]().PutAll(m))
	assert.Equal(t, fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"}), fp.MapOf(map[int]string{1: "s1"}).PutAll(m))
}

func TestMapEquals(t *testing.T) {
	assert.True(t, fp.EmptyMap[int, string]().Equals(fp.MapOf(map[int]string{})))
	assert.True(t, fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"}).Equals(fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"})))
	assert.False(t, fp.MapOf(map[int]string{1: "s1"}).Equals(fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"})))
}

func TestMapSize(t *testing.T) {
	assert.Equal(t, 0, fp.EmptyMap[int, string]().Size())
	assert.Equal(t, 3, fp.MapOf(map[int]string{1: "s1", 2: "s2", 3: "s3"}).Size())
}

func TestMapIsEmpty(t *testing.T) {
	assert.True(t, fp.EmptyMap[int, string]().IsEmpty())
	assert.False(t, fp.MapOf(map[int]string{1: "s1"}).IsEmpty())
}

func TestMapNonEmpty(t *testing.T) {
	assert.False(t, fp.EmptyMap[int, string]().NonEmpty())
	assert.True(t, fp.MapOf(map[int]string{1: "s1"}).NonEmpty())
}

func TestMapGet(t *testing.T) {
	assert.Equal(t, fp.SomeOf("s2"), fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Get(2))
	assert.True(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Get(3).NonDefined())
}

func TestMapGetOrElse(t *testing.T) {
	assert.Equal(t, "s2", fp.MapOf(map[int]string{1: "s1", 2: "s2"}).GetOrElse(2, "s"))
	assert.Equal(t, "s", fp.MapOf(map[int]string{1: "s1", 2: "s2"}).GetOrElse(3, "s"))
}

func TestMapRemove(t *testing.T) {
	assert.Equal(t, fp.MapOf(map[int]string{}), fp.MapOf(map[int]string{}).Remove(0))
	assert.Equal(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}), fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Remove(0))
	assert.Equal(t, fp.MapOf(map[int]string{2: "s2"}), fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Remove(1))
}

func TestMapContainsKey(t *testing.T) {
	assert.False(t, fp.MapOf(map[int]string{}).ContainsKey(0))
	assert.False(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}).ContainsKey(0))
	assert.True(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}).ContainsKey(1))
}

func TestMapContainsValue(t *testing.T) {
	assert.False(t, fp.MapOf(map[int]string{}).ContainsValue("s1"))
	assert.False(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}).ContainsValue("s3"))
	assert.True(t, fp.MapOf(map[int]string{1: "s1", 2: "s2"}).ContainsValue("s2"))
}

func TestMapEntries(t *testing.T) {
	assert.True(t, fp.MapOf(map[int]string{}).Entries().IsEmpty())
	assert.True(t, setEqual(fp.ArrayOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")), fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Entries()))
}

func TestMapKeys(t *testing.T) {
	assert.True(t, fp.MapOf(map[int]string{}).Keys().IsEmpty())
	assert.True(t, setEqual(fp.ArrayOf(1, 2), fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Keys()))
}

func TestMapValues(t *testing.T) {
	assert.True(t, fp.MapOf(map[int]string{}).Values().IsEmpty())
	assert.True(t, setEqual(fp.ArrayOf("s1", "s2"), fp.MapOf(map[int]string{1: "s1", 2: "s2"}).Values()))
}

func setEqual[T any](sa fp.Seq[T], sb fp.Seq[T]) bool {
	if sa.Size() != sb.Size() {
		return false
	}
	return sa.ForAll(func(x T) bool { return sb.ContainsElement(x) })
}
