package redblackbst

import (
	"testing"
)

type K string

func (k K) Compare(other KType) int {
	ks := string(k)
	os := string(other.(K))
	if ks < os {
		return -1
	}
	if ks > os {
		return 1
	}
	return 0
}

func putEmpty(t *testing.T, tree *RedBlack, k KType, v VType) {
	old, overwrite := tree.Put(k, v)
	if overwrite {
		t.Fatalf("shouldnt have overwritten, old %#v: %#v", old, tree)
	}
}

func putOver(t *testing.T, tree *RedBlack, k KType, v VType, replace VType) {
	old, overwrite := tree.Put(k, v)
	if !overwrite {
		t.Fatalf("should have overwritten %v, old %#v: %#v", replace, old, tree)
	}
	if old != replace {
		t.Fatalf("want old %#v, got %#v", replace, old)
	}
}

func getCheckEmpty(t *testing.T, tree *RedBlack, k KType) {
	got, ok := tree.Get(k)
	if ok {
		t.Fatalf("should have got nothing, got %#v", got)
	}
}

func getCheckVal(t *testing.T, tree *RedBlack, k KType, want VType) {
	gotv, ok := tree.Get(k)
	if !ok {
		t.Fatalf("want %q, got nothing: %#v", want, tree)
	}
	got := gotv.(string)
	if want != got {
		t.Fatalf("want %q got %q", want, got)
	}
}

func TestTreeEmpty(t *testing.T) {
	tree := New()
	if !tree.IsEmpty() {
		t.Fatalf("tree is not empty: %#v", tree)
	}

	putEmpty(t, tree, K("hello"), "world")

	if tree.IsEmpty() {
		t.Fatalf("tree is empty: %#v", tree)
	}
	t.Logf("tree=%#v", tree)
	if !tree.Delete(K("hello")) {
		t.Errorf("couldn't delete key from tree: %#v", tree)
	}
	if !tree.IsEmpty() {
		t.Fatalf("tree is not empty: %#v", tree)
	}
}

func TestTreeSize(t *testing.T) {
	tree := New()
	if tree.Size() != 0 {
		t.Fatalf("tree is not empty: %#v", tree)
	}
	putEmpty(t, tree, K("hello"), "world")

	tree.Clear()
	if tree.Size() != 0 {
		t.Fatalf("tree is not empty: %#v", tree)
	}
}

func TestCanPutAndGet(t *testing.T) {
	tree := New()
	k := K("hello")
	want := "world"

	for i := 0; i < 100; i++ {
		putEmpty(t, tree, k, want)
		getCheckVal(t, tree, k, want)
		tree.Clear()
	}
}

func TestCanDeleteSomethingAbsent(t *testing.T) {
	tree := New()

	k := K("im not there")
	if tree.Delete(k) {
		t.Fatalf("shouldn't be able to delete %q", k)
	}

	putEmpty(t, tree, K("im not there lol"), "hello")
}

func TestCanPutAndGetManyElements(t *testing.T) {

	kv := makeKV()
	tree := New()

	i := 0
	for k, v := range kv {
		i++

		putEmpty(t, tree, k, v)

		wantSize := i
		gotSize := tree.Size()
		if wantSize != gotSize {
			t.Errorf("want size %d, got %d", wantSize, gotSize)
		}

		getCheckVal(t, tree, k, v)
	}

	wantSize := len(kv)
	gotSize := tree.Size()
	if wantSize != gotSize {
		t.Errorf("want size %d, got %d", wantSize, gotSize)
	}

	for k, v := range kv {
		getCheckVal(t, tree, k, v)
	}

}

func TestCanPutAndGetAndOverwrite(t *testing.T) {
	tree := New()
	k := K("hello")
	want := "world"
	nextWant := "le monde"

	for i := 0; i < 100; i++ {
		putEmpty(t, tree, k, want)

		getCheckVal(t, tree, k, want)

		putOver(t, tree, k, nextWant, want)

		getCheckVal(t, tree, k, nextWant)

		tree.Clear()
	}
}

func TestCanPutAndDelete(t *testing.T) {
	tree := New()
	k := K("hello")
	want := "world"

	for i := 0; i < 100; i++ {

		if tree.Size() != 0 {
			t.Fatalf("tree is not empty: %#v", tree)
		}
		if !tree.IsEmpty() {
			t.Fatalf("tree is not empty: %#v", tree)
		}

		putEmpty(t, tree, k, want)

		if tree.Size() != 1 {
			t.Fatalf("tree is  empty: %#v", tree)
		}
		if tree.IsEmpty() {
			t.Fatalf("tree is  empty: %#v", tree)
		}

		if !tree.Delete(k) {
			t.Errorf("couldn't delete key from tree: %#v", tree)
		}
	}

	if tree.Size() != 0 {
		t.Fatalf("tree is not empty: %#v", tree)
	}
	if !tree.IsEmpty() {
		t.Fatalf("tree is not empty: %#v", tree)
	}
}

func TestPanicDelete(t *testing.T) {
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

	printTreeStats(tree, "start")

	tree.Delete(K("key:abature"))
	tree.Delete(K("key:abbacy"))
	tree.Delete(K("key:abasia"))
	tree.Delete(K("key:abask"))
	tree.Delete(K("key:abator"))
	tree.Delete(K("key:abaton"))
	tree.Delete(K("key:abaze"))
	tree.Delete(K("key:abastardize"))

	printTreeStats(tree, "before-before-panic")

	tree.Delete(K("key:abatement"))

	printTreeStats(tree, "before-panic")

	defer func() {
		printTreeStats(tree, "after-panic")
	}()

	tree.Delete(K("key:abater"))
}
