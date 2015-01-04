package redblackbst

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"
)

func putEmpty(t *testing.T, tree *RedBlack, k KType) {
	overwrite := tree.Put(k)
	if overwrite {
		t.Fatalf("shouldnt have overwritten, %#v", tree)
	}
}

func putOver(t *testing.T, tree *RedBlack, k KType) {
	overwrite := tree.Put(k)
	if !overwrite {
		t.Fatalf("should have overwritten %v, %#v", k, tree)
	}
}

func checkNotContains(t *testing.T, tree *RedBlack, k KType) {
	ok := tree.Contains(k)
	if ok {
		t.Fatalf("should have got nothing for %#v: %#v", k, tree)
	}
}

func checkContains(t *testing.T, tree *RedBlack, k KType) {
	ok := tree.Contains(k)
	if !ok {
		t.Fatalf("want %#v, got nothing: %#v", k, tree)
	}
}

func TestTreeEmpty(t *testing.T) {
	tree := NewRedBlack()
	if !tree.IsEmpty() {
		t.Fatalf("tree is not empty: %#v", tree)
	}

	putEmpty(t, tree, K("hello"))

	if tree.IsEmpty() {
		t.Fatalf("tree is empty: %#v", tree)
	}

	if !tree.Delete(K("hello")) {
		t.Errorf("couldn't delete key from tree: %#v", tree)
	}
	if !tree.IsEmpty() {
		t.Fatalf("tree is not empty: %#v", tree)
	}
}

func TestTreeSize(t *testing.T) {
	tree := NewRedBlack()
	if tree.Size() != 0 {
		t.Fatalf("tree is not empty: %#v", tree)
	}
	putEmpty(t, tree, K("hello"))

	tree.Clear()
	if tree.Size() != 0 {
		t.Fatalf("tree is not empty: %#v", tree)
	}
}

func TestCanPutAndGet(t *testing.T) {
	tree := NewRedBlack()
	k := K("hello")

	for i := 0; i < 100; i++ {
		putEmpty(t, tree, k)
		checkContains(t, tree, k)
		tree.Clear()
	}
}

func TestCanDeleteSomethingAbsent(t *testing.T) {
	tree := NewRedBlack()

	k := K("im not there")
	if tree.Delete(k) {
		t.Fatalf("shouldn't be able to delete %q", k)
	}

	putEmpty(t, tree, K("im not there lol"))
}

func TestCanPutAndGetManyElements(t *testing.T) {

	kv := makeKV()
	tree := NewRedBlack()

	i := 0
	for k := range kv {
		i++

		putEmpty(t, tree, k)

		wantSize := i
		gotSize := tree.Size()
		if wantSize != gotSize {
			t.Errorf("want size %d, got %d", wantSize, gotSize)
		}

		checkContains(t, tree, k)
	}

	wantSize := len(kv)
	gotSize := tree.Size()
	if wantSize != gotSize {
		t.Errorf("want size %d, got %d", wantSize, gotSize)
	}

	for k := range kv {
		checkContains(t, tree, k)
	}

	for k := range kv {

		if !tree.Delete(k) {
			t.Fatalf("should have deleted %#v", k)
		}
		checkNotContains(t, tree, k)
	}

}

func TestCanPutAndGetAndOverwrite(t *testing.T) {
	tree := NewRedBlack()
	k := K("hello")

	for i := 0; i < 100; i++ {
		putEmpty(t, tree, k)

		checkContains(t, tree, k)

		putOver(t, tree, k)

		checkContains(t, tree, k)

		tree.Clear()
	}
}

func TestCanPutAndDelete(t *testing.T) {
	tree := NewRedBlack()
	k := K("hello")

	for i := 0; i < 100; i++ {

		if tree.Size() != 0 {
			t.Fatalf("tree is not empty: %#v", tree)
		}
		if !tree.IsEmpty() {
			t.Fatalf("tree is not empty: %#v", tree)
		}

		putEmpty(t, tree, k)

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

func TestEmptyTreeHasNoMin(t *testing.T) {
	tree := NewRedBlack()
	k, ok := tree.Min()
	if ok {
		t.Fatalf("should have no min, got %v", k)
	}
}

func TestEmptyTreeCorrectMin(t *testing.T) {
	tree := NewRedBlack()

	perm := rand.Perm(100)

	for _, i := range perm {
		tree.Put(Int(i))
	}
	k, ok := tree.Min()
	if !ok {
		t.Fatalf("should have max, got nothing")
	}
	if k != Int(0) {
		t.Errorf("want k %d, got %v", 0, k)
	}
}

func TestEmptyTreeHasNoMax(t *testing.T) {
	tree := NewRedBlack()
	k, ok := tree.Max()
	if ok {
		t.Fatalf("should have no max, got %v", k)
	}
}

func TestEmptyTreeCorrectMax(t *testing.T) {
	tree := NewRedBlack()

	perm := rand.Perm(100)

	for _, i := range perm {
		tree.Put(Int(i))
	}
	k, ok := tree.Max()
	if !ok {
		t.Fatalf("should have max, got nothing")
	}
	if k != Int(99) {
		t.Errorf("want k %d, got %v", 99, k)
	}
}

func TestEmptyTreeHasMaxOneElement(t *testing.T) {
	tree := NewRedBlack()

	tree.Put(Int(100))
	tree.Put(Int(98))
	tree.Put(Int(99))
	tree.Put(Int(101))

	k, ok := tree.Max()
	if !ok {
		t.Fatalf("should have max, got nothing")
	}
	if k != Int(101) {
		t.Errorf("want k %d, got %v", 101, k)
	}
}

func TestGetCorrectFloor(t *testing.T) {
	tree := NewRedBlack()
	k, ok := tree.Floor(Int(0))
	if ok {
		t.Errorf("should not have found floor, got %v", k)
	}

	tree.Put(Int(1))

	k, ok = tree.Floor(Int(0))
	if ok {
		t.Errorf("should not have found floor, got %v", k)
	}

	k, ok = tree.Floor(Int(2))
	if !ok {
		t.Errorf("should have found floor")
	}
	if k != Int(1) {
		t.Errorf("want k %v, got k %v", Int(1), k)
	}

	for i := 0; i < 100; i++ {
		if i == 50 {
			continue
		}
		tree.Put(Int(i))
	}
	k, ok = tree.Floor(Int(50))
	if !ok {
		t.Errorf("should not have found floor")
	}
	if k != Int(49) {
		t.Errorf("want k %v, got k %v", Int(49), k)
	}

	for i := 100; i < 200; i++ {
		tree.Put(Int(i))
	}
	k, ok = tree.Floor(Int(150))
	if !ok {
		t.Errorf("should not have found floor")
	}
	if k != Int(150) {
		t.Errorf("want k %v, got k %v", Int(150), k)
	}
}

func TestGetCorrectCeiling(t *testing.T) {
	tree := NewRedBlack()
	k, ok := tree.Ceiling(Int(0))
	if ok {
		t.Errorf("should not have found ceiling, got %v", k)
	}

	tree.Put(Int(1))

	k, ok = tree.Ceiling(Int(2))
	if ok {
		t.Errorf("should not have found ceiling, got %v", k)
	}

	k, ok = tree.Ceiling(Int(0))
	if !ok {
		t.Errorf("should have found ceiling")
	}
	if k != Int(1) {
		t.Errorf("want k %v, got k %v", Int(1), k)
	}

	for i := 0; i < 100; i++ {
		if i == 50 {
			continue
		}
		tree.Put(Int(i))
	}
	k, ok = tree.Ceiling(Int(50))
	if !ok {
		t.Errorf("should not have found ceiling")
	}
	if k != Int(51) {
		t.Errorf("want k %v, got k %v", Int(51), k)
	}

	for i := 100; i < 200; i++ {
		tree.Put(Int(i))
	}
	k, ok = tree.Ceiling(Int(150))
	if !ok {
		t.Errorf("should not have found ceiling")
	}
	if k != Int(150) {
		t.Errorf("want k %v, got k %v", Int(150), k)
	}
}

func TestGetCorrectSelect(t *testing.T) {

	tree := NewRedBlack()

	gotk, ok := tree.Select(0)
	if ok {
		t.Errorf("should not be able to select, got %v", gotk)
	}

	for i := 0; i < 100; i++ {
		wantk := Int(i)
		putEmpty(t, tree, wantk)

		gotk, ok = tree.Select(i)
		if !ok {
			t.Errorf("should be able to select")
		}

		if !reflect.DeepEqual(gotk, wantk) {
			t.Errorf("want k %v, got %v", wantk, gotk)
		}
	}
	for i := 0; i < 100; i++ {
		wantk := Int(i)

		gotk, ok = tree.Select(i)
		if !ok {
			t.Errorf("should be able to select")
		}

		if !reflect.DeepEqual(gotk, wantk) {
			t.Errorf("want k %v, got %v", wantk, gotk)
		}
	}
}

func TestGetCorrectRank(t *testing.T) {

	tree := NewRedBlack()

	if tree.Rank(Int(0)) != 0 {
		t.Errorf("should not have a rank, got %v", tree.Rank(Int(0)))
	}

	for i := 0; i < 100; i++ {
		wantk := Int(i)
		putEmpty(t, tree, wantk)

		goti := tree.Rank(wantk)
		if goti != i {
			t.Errorf("got rank %d, want %d", goti, i)
		}
	}

	for i := 100; i > 0; i-- {
		wantk := Int(i)

		goti := tree.Rank(wantk)
		if goti != i {
			t.Fatalf("got rank %d, want %d", goti, i)
		}
	}
}

func TestCanDeleteMin(t *testing.T) {
	tree := NewRedBlack()
	kv := map[Int]Int{}
	want := []Int{}
	for i := 0; i < 100; i++ {
		kv[Int(i)] = Int(i)
		want = append(want, Int(i))
	}

	for k := range kv {
		tree.Put(k)
	}

	got := []Int{}
	for {
		k, ok := tree.DeleteMin()
		if !ok {
			break
		}
		got = append(got, k.(Int))
	}

	if !reflect.DeepEqual(want, got) {
		t.Logf("want=%#v", want)
		t.Logf(" got=%#v", got)
		t.Errorf("mismatch!")
	}
}

func TestCanDeleteMax(t *testing.T) {
	tree := NewRedBlack()
	kv := map[Int]Int{}
	want := []Int{}
	for i := 100; i > 0; i-- {
		kv[Int(i)] = Int(i)
		want = append(want, Int(i))
	}

	for k := range kv {
		tree.Put(k)
	}

	got := []Int{}
	for {
		k, ok := tree.DeleteMax()
		if !ok {
			if len(got) != len(want) {
				t.Errorf("k,v,ok = %v,%v", k, ok)
			}
			break
		}
		got = append(got, k.(Int))
	}

	if !reflect.DeepEqual(want, got) {
		t.Logf("want=%#v", want)
		t.Logf(" got=%#v", got)
		t.Errorf("mismatch!")
	}
}

func TestCanVisitAllKeysWhenPutBackward(t *testing.T) {
	tree := NewRedBlack()
	want := make([]Int, 100)
	for i := 99; i >= 0; i-- {
		tree.Put(Int(i))
		want[i] = Int(i)
	}

	got := []Int{}
	tree.Keys(func(k KType) bool {
		got = append(got, k.(Int))
		return true
	})

	if !reflect.DeepEqual(want, got) {
		t.Logf("want=%#v", want)
		t.Logf(" got=%#v", got)
		t.Errorf("mismatch!")
	}
}

func TestCanVisitAllKeys(t *testing.T) {
	tree := NewRedBlack()
	kv := map[Int]Int{}
	want := []Int{}
	for i := 0; i < 100; i++ {
		kv[Int(i)] = Int(i)
		want = append(want, Int(i))
	}

	for k := range kv {
		tree.Put(k)
	}

	got := []Int{}
	tree.Keys(func(k KType) bool {
		got = append(got, k.(Int))
		return true
	})

	if !reflect.DeepEqual(want, got) {
		t.Logf("want=%#v", want)
		t.Logf(" got=%#v", got)
		t.Errorf("mismatch!")
	}
}

func TestCanAbortVisitingAllKeys(t *testing.T) {
	tree := NewRedBlack()
	want := []Int{}
	for i := 0; i < 100; i++ {
		want = append(want, Int(i))
		tree.Put(Int(i))
	}

	want = want[:50]

	got := []Int{}
	tree.Keys(func(k KType) bool {
		got = append(got, k.(Int))
		return len(got) < 50
	})

	if !reflect.DeepEqual(want, got) {
		t.Logf("want=%#v", want)
		t.Logf(" got=%#v", got)
		t.Errorf("mismatch!")
	}
}

func TestCanVisitEmptyTree(t *testing.T) {
	tree := NewRedBlack()
	tree.Keys(func(k KType) bool {
		t.Errorf("shouldn not be called, got %v", k)
		return true
	})

}

func TestCanExportToDot(t *testing.T) {
	tree := NewRedBlack()
	for i := 0; i < 100; i++ {
		tree.Put(Int(i))
	}
	buf := bytes.NewBuffer(nil)
	_, err := tree.DotGraph(buf, "test-graph")
	if err != nil {
		t.Fatal(err)
	}
	openDot(buf)
}
