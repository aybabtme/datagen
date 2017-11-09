package redblackbst

import (
	"crypto/rand"
	mrand "math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/oklog/ulid"
)

func TestRegressionPanicDelete(t *testing.T) {
	tree := NewRedBlack()
	tree.Put(K("key:abature"), "value:abature")
	tree.Put(K("key:abbacy"), "value:abbacy")
	tree.Put(K("key:abasia"), "value:abasia")
	tree.Put(K("key:abask"), "value:abask")
	tree.Put(K("key:abator"), "value:abator")
	tree.Put(K("key:abaton"), "value:abaton")
	tree.Put(K("key:abaze"), "value:abaze")
	tree.Put(K("key:abastardize"), "value:abastardize")
	tree.Put(K("key:abatement"), "value:abatement")
	tree.Put(K("key:abater"), "value:abater")
	tree.Put(K("key:abatis"), "value:abatis")
	tree.Put(K("key:abash"), "value:abash")
	tree.Put(K("key:abate"), "value:abate")
	tree.Put(K("key:abatised"), "value:abatised")
	tree.Put(K("key:Abassin"), "value:Abassin")
	tree.Put(K("key:abb"), "value:abb")
	tree.Put(K("key:abatable"), "value:abatable")
	tree.Put(K("key:abaxile"), "value:abaxile")

	tree.Delete(K("key:abature"))
	tree.Delete(K("key:abbacy"))
	tree.Delete(K("key:abasia"))
	tree.Delete(K("key:abask"))
	tree.Delete(K("key:abator"))
	tree.Delete(K("key:abaton"))
	tree.Delete(K("key:abaze"))
	tree.Delete(K("key:abastardize"))
	tree.Delete(K("key:abatement"))
	tree.Delete(K("key:abater"))
}

func TestLexicalStringOrder(t *testing.T) {
	samples := 100
	minSize := 2
	maxSize := 20
	permCount := 100

	generateKeys := func(n int) []string {
		keys := make([]string, 0, n)
		for i := 0; i < n; i++ {
			keys = append(keys, ulid.MustNew(ulid.Now(), rand.Reader).String())
		}
		return keys
	}

	for i := 0; i < samples; i++ {
		want := generateKeys(minSize + mrand.Intn(maxSize-minSize))
		for j := 0; j < permCount; j++ {
			tree := NewRedBlack()

			for _, idx := range mrand.Perm(len(want)) {
				k := want[idx]
				tree.Put(K(k), "value:"+k)
			}
			var got []string
			tree.Keys(func(k KType, _ VType) bool {
				got = append(got, string(k.(K)))
				return true
			})
			sort.Strings(want)
			if !reflect.DeepEqual(want, got) {
				t.Logf("want=%#v", want)
				t.Logf(" got=%#v", got)
				t.Fatalf("mismatch!")
			}
		}
	}
}
