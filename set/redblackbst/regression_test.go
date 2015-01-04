package redblackbst

import (
	"testing"
)

func TestRegressionPanicDelete(t *testing.T) {
	tree := NewRedBlack()
	tree.Put(K("key:abature"))
	tree.Put(K("key:abbacy"))
	tree.Put(K("key:abasia"))
	tree.Put(K("key:abask"))
	tree.Put(K("key:abator"))
	tree.Put(K("key:abaton"))
	tree.Put(K("key:abaze"))
	tree.Put(K("key:abastardize"))
	tree.Put(K("key:abatement"))
	tree.Put(K("key:abater"))
	tree.Put(K("key:abatis"))
	tree.Put(K("key:abash"))
	tree.Put(K("key:abate"))
	tree.Put(K("key:abatised"))
	tree.Put(K("key:Abassin"))
	tree.Put(K("key:abb"))
	tree.Put(K("key:abatable"))
	tree.Put(K("key:abaxile"))

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
