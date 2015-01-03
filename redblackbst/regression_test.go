package redblackbst

import (
	"testing"
)

func TestRegressionPanicDelete(t *testing.T) {
	tree := New()
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
