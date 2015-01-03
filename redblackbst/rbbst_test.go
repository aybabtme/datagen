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

func TestTreeEmpty(t *testing.T) {
	tree := New()
	if !tree.IsEmpty() {
		t.Fatalf("tree is not empty: %#v", tree)
	}
	tree.Put(K("hello"), "world")
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
	tree.Put(K("hello"), "world")
	if tree.Size() != 1 {
		t.Fatalf("tree is empty: %#v", tree)
	}

	tree.Clear()
	if tree.Size() != 0 {
		t.Fatalf("tree is not empty: %#v", tree)
	}
}

func TestCanPutAndGet(t *testing.T) {
	tree := New()
	k := K("hello")
	want := "world"
	tree.Put(k, want)

	for i := 0; i < 100; i++ {
		tree.Put(k, want)

		gotv, ok := tree.Get(k)
		if !ok {
			t.Fatalf("want %q, got nothing: %#v", want, tree)
		}
		got := gotv.(string)
		if want != got {
			t.Fatalf("want %q got %q", want, got)
		}
		tree.Clear()
	}
}

func TestCanPutAndGetAndOverwrite(t *testing.T) {
	tree := New()
	k := K("hello")
	want := "world"
	nextWant := "le monde"
	tree.Put(k, want)

	for i := 0; i < 100; i++ {
		tree.Put(k, want)

		gotv, ok := tree.Get(k)
		if !ok {
			t.Fatalf("want %q, got nothing: %#v", want, tree)
		}
		got := gotv.(string)
		if want != got {
			t.Fatalf("want %q got %q", want, got)
		}

		tree.Put(k, nextWant)

		gotv, ok = tree.Get(k)
		if !ok {
			t.Fatalf("nextWant %q, got nothing: %#v", nextWant, tree)
		}
		got = gotv.(string)
		if nextWant != got {
			t.Fatalf("nextWant %q got %q", nextWant, got)
		}

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

		tree.Put(k, want)

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
