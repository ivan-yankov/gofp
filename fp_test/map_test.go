package fp_test

import (
	"testing"

	"github.com/ivan-yankov/gofp/fp"
	"github.com/stretchr/testify/assert"
)

func TestMapPut(t *testing.T) {
	assert.True(t, fp.MapOf[int, string]().Put(1, "s1").Equals(fp.MapOf(fp.PairOf(1, "s1"))))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1")).Put(1, "s2").Equals(fp.MapOf(fp.PairOf(1, "s2"))))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1")).Put(2, "s2").Equals(fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"))))
}

func TestMapPutAll(t *testing.T) {
	m := fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3"))
	assert.True(t, fp.MapOf[int, string]().PutAll(m).Equals(fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3"))))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s5")).PutAll(m).Equals(fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3"))))
}

func TestMapEquals(t *testing.T) {
	assert.True(t, fp.MapOf[int, string]().Equals(fp.MapOf[int, string]()))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3")).Equals(fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3"))))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3")).Equals(fp.MapOf(fp.PairOf(3, "s3"), fp.PairOf(2, "s2"), fp.PairOf(1, "s1"))))
	assert.False(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3")).Equals(fp.MapOf(fp.PairOf(1, "s3"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3"))))
	assert.False(t, fp.MapOf(fp.PairOf(1, "s1")).Equals(fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3"))))
}

func TestMapSize(t *testing.T) {
	assert.Equal(t, 0, fp.MapOf[int, string]().Size())
	assert.Equal(t, 3, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2"), fp.PairOf(3, "s3")).Size())
}

func TestMapIsEmpty(t *testing.T) {
	assert.True(t, fp.MapOf[int, string]().IsEmpty())
	assert.False(t, fp.MapOf(fp.PairOf(1, "s1")).IsEmpty())
}

func TestMapNonEmpty(t *testing.T) {
	assert.False(t, fp.MapOf[int, string]().NonEmpty())
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1")).NonEmpty())
}

func TestMapGet(t *testing.T) {
	assert.Equal(t, fp.SomeOf("s2"), fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Get(2))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Get(3).NonDefined())
}

func TestMapGetOrElse(t *testing.T) {
	assert.Equal(t, "s2", fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).GetOrElse(2, "s"))
	assert.Equal(t, "s", fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).GetOrElse(3, "s"))
}

func TestMapRemove(t *testing.T) {
	assert.Equal(t, fp.MapOf[int, string](), fp.MapOf[int, string]().Remove(0))
	assert.Equal(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")), fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Remove(0))
	assert.Equal(t, fp.MapOf(fp.PairOf(2, "s2")), fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Remove(1))
}

func TestMapContainsKey(t *testing.T) {
	assert.False(t, fp.MapOf[int, string]().ContainsKey(0))
	assert.False(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).ContainsKey(0))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).ContainsKey(1))
}

func TestMapContainsValue(t *testing.T) {
	assert.False(t, fp.MapOf[int, string]().ContainsValue("s1"))
	assert.False(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).ContainsValue("s3"))
	assert.True(t, fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).ContainsValue("s2"))
}

func TestMapEntries(t *testing.T) {
	assert.True(t, fp.MapOf[int, string]().Entries().IsEmpty())
	assert.True(t, fp.SeqSetEquals(fp.ArrayOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")), fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Entries()))
}

func TestMapKeys(t *testing.T) {
	assert.True(t, fp.MapOf[int, string]().Keys().IsEmpty())
	assert.True(t, fp.SeqSetEquals(fp.ArrayOf(1, 2), fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Keys()))
}

func TestMapValues(t *testing.T) {
	assert.True(t, fp.MapOf[int, string]().Values().IsEmpty())
	assert.True(t, fp.SeqSetEquals(fp.ArrayOf("s1", "s2"), fp.MapOf(fp.PairOf(1, "s1"), fp.PairOf(2, "s2")).Values()))
}
