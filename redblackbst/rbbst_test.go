package redblackbst

import (
	"math/rand"
	"reflect"
	"testing"
)

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

	if _, ok := tree.Delete(K("hello")); !ok {
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
	if _, ok := tree.Delete(k); ok {
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

	for k, v := range kv {
		old, ok := tree.Delete(k)
		if !ok {
			t.Fatalf("should have deleted %#v", k)
		}

		if old.(string) != v {
			t.Fatalf("want %q got %q", v, old.(string))
		}
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

		if _, ok := tree.Delete(k); !ok {
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
	tree := New()
	k, v, ok := tree.Min()
	if ok {
		t.Fatalf("should have no min, got %v-%v", k, v)
	}
}

func TestEmptyTreeCorrectMin(t *testing.T) {
	tree := New()

	perm := rand.Perm(100)

	for _, i := range perm {
		tree.Put(Int(i), Int(i))
	}
	k, v, ok := tree.Min()
	if !ok {
		t.Fatalf("should have max, got nothing")
	}
	if k != Int(0) {
		t.Errorf("want k %d, got %v", 0, k)
	}
	if v != Int(0) {
		t.Errorf("want k %d, got %v", 0, k)
	}
}

func TestEmptyTreeHasNoMax(t *testing.T) {
	tree := New()
	k, v, ok := tree.Max()
	if ok {
		t.Fatalf("should have no max, got %v-%v", k, v)
	}
}

func TestEmptyTreeCorrectMax(t *testing.T) {
	tree := New()

	perm := rand.Perm(100)

	for _, i := range perm {
		tree.Put(Int(i), Int(i))
	}
	k, v, ok := tree.Max()
	if !ok {
		t.Fatalf("should have max, got nothing")
	}
	if k != Int(99) {
		t.Errorf("want k %d, got %v", 99, k)
	}
	if v != Int(99) {
		t.Errorf("want k %d, got %v", 99, k)
	}
}

func TestEmptyTreeHasMaxOneElement(t *testing.T) {
	tree := New()

	tree.Put(Int(100), Int(100))
	tree.Put(Int(98), Int(98))
	tree.Put(Int(99), Int(99))
	tree.Put(Int(101), Int(101))

	k, v, ok := tree.Max()
	if !ok {
		t.Fatalf("should have max, got nothing")
	}
	if k != Int(101) {
		t.Errorf("want k %d, got %v", 101, k)
	}
	if v != Int(101) {
		t.Errorf("want k %d, got %v", 101, k)
	}
}

func TestGetCorrectFloor(t *testing.T) {
	tree := New()
	k, v, ok := tree.Floor(Int(0))
	if ok {
		t.Errorf("should not have found floor, got %v->%v", k, v)
	}

	tree.Put(Int(1), Int(1))

	k, v, ok = tree.Floor(Int(0))
	if ok {
		t.Errorf("should not have found floor, got %v->%v", k, v)
	}

	k, v, ok = tree.Floor(Int(2))
	if !ok {
		t.Errorf("should have found floor")
	}
	if k != Int(1) {
		t.Errorf("want k %v, got k %v, v %v", Int(1), k, v)
	}

	for i := 0; i < 100; i++ {
		if i == 50 {
			continue
		}
		tree.Put(Int(i), Int(i))
	}
	k, v, ok = tree.Floor(Int(50))
	if !ok {
		t.Errorf("should not have found floor")
	}
	if k != Int(49) {
		t.Errorf("want k %v, got k %v, v %v", Int(49), k, v)
	}

	for i := 100; i < 200; i++ {
		tree.Put(Int(i), Int(i))
	}
	k, v, ok = tree.Floor(Int(150))
	if !ok {
		t.Errorf("should not have found floor")
	}
	if k != Int(150) {
		t.Errorf("want k %v, got k %v, v %v", Int(150), k, v)
	}
}

func TestGetCorrectCeiling(t *testing.T) {
	tree := New()
	k, v, ok := tree.Ceiling(Int(0))
	if ok {
		t.Errorf("should not have found ceiling, got %v->%v", k, v)
	}

	tree.Put(Int(1), Int(1))

	k, v, ok = tree.Ceiling(Int(2))
	if ok {
		t.Errorf("should not have found ceiling, got %v->%v", k, v)
	}

	k, v, ok = tree.Ceiling(Int(0))
	if !ok {
		t.Errorf("should have found ceiling")
	}
	if k != Int(1) {
		t.Errorf("want k %v, got k %v, v %v", Int(1), k, v)
	}

	for i := 0; i < 100; i++ {
		if i == 50 {
			continue
		}
		tree.Put(Int(i), Int(i))
	}
	k, v, ok = tree.Ceiling(Int(50))
	if !ok {
		t.Errorf("should not have found ceiling")
	}
	if k != Int(51) {
		t.Errorf("want k %v, got k %v, v %v", Int(51), k, v)
	}

	for i := 100; i < 200; i++ {
		tree.Put(Int(i), Int(i))
	}
	k, v, ok = tree.Ceiling(Int(150))
	if !ok {
		t.Errorf("should not have found ceiling")
	}
	if k != Int(150) {
		t.Errorf("want k %v, got k %v, v %v", Int(150), k, v)
	}
}

func TestGetCorrectSelect(t *testing.T) {

	tree := New()

	gotk, gotv, ok := tree.Select(0)
	if ok {
		t.Errorf("should not be able to select, got %v->%v", gotk, gotv)
	}

	for i := 0; i < 100; i++ {
		wantk, wantv := Int(i), Int(i)
		putEmpty(t, tree, wantk, wantv)

		gotk, gotv, ok = tree.Select(i)
		if !ok {
			t.Errorf("should be able to select")
		}

		if !reflect.DeepEqual(gotk, wantk) {
			t.Errorf("want k %v, got %v", wantk, gotk)
		}
		if !reflect.DeepEqual(gotv, wantv) {
			t.Errorf("want v %v, got %v", wantv, gotv)
		}
	}
	for i := 0; i < 100; i++ {
		wantk, wantv := Int(i), Int(i)

		gotk, gotv, ok = tree.Select(i)
		if !ok {
			t.Errorf("should be able to select")
		}

		if !reflect.DeepEqual(gotk, wantk) {
			t.Errorf("want k %v, got %v", wantk, gotk)
		}
		if !reflect.DeepEqual(gotv, wantv) {
			t.Errorf("want v %v, got %v", wantv, gotv)
		}
	}
}

func TestGetCorrectRank(t *testing.T) {

	tree := New()

	if tree.Rank(Int(0)) != 0 {
		t.Errorf("should not have a rank, got %v", tree.Rank(Int(0)))
	}

	for i := 0; i < 100; i++ {
		wantk, wantv := Int(i), Int(i)
		putEmpty(t, tree, wantk, wantv)

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
	tree := New()
	kv := map[Int]Int{}
	want := []Int{}
	for i := 0; i < 100; i++ {
		kv[Int(i)] = Int(i)
		want = append(want, Int(i))
	}

	for k, v := range kv {
		tree.Put(k, v)
	}

	got := []Int{}
	for {
		k, _, ok := tree.DeleteMin()
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
	tree := New()
	kv := map[Int]Int{}
	want := []Int{}
	for i := 100; i > 0; i-- {
		kv[Int(i)] = Int(i)
		want = append(want, Int(i))
	}

	for k, v := range kv {
		tree.Put(k, v)
	}

	got := []Int{}
	for {
		k, v, ok := tree.DeleteMax()
		if !ok {
			if len(got) != len(want) {
				t.Errorf("k,v,ok = %v,%v,%v", k, v, ok)
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
	tree := New()
	want := make([]Int, 100)
	for i := 99; i >= 0; i-- {
		tree.Put(Int(i), Int(i))
		want[i] = Int(i)
	}

	got := []Int{}
	tree.Keys(func(k KType, v VType) bool {
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
	tree := New()
	kv := map[Int]Int{}
	want := []Int{}
	for i := 0; i < 100; i++ {
		kv[Int(i)] = Int(i)
		want = append(want, Int(i))
	}

	for k, v := range kv {
		tree.Put(k, v)
	}

	got := []Int{}
	tree.Keys(func(k KType, v VType) bool {
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
	tree := New()
	want := []Int{}
	for i := 0; i < 100; i++ {
		want = append(want, Int(i))
		tree.Put(Int(i), Int(i))
	}

	want = want[:50]

	got := []Int{}
	tree.Keys(func(k KType, v VType) bool {
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
	tree := New()
	tree.Keys(func(k KType, v VType) bool {
		t.Errorf("shouldn not be called, got %v->%v", k, v)
		return true
	})

}
