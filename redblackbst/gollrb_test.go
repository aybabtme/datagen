// Adapted from github.com/petar/GoLLRB/llrb/llrb_test.go

// Copyright 2010 Petar Maymounkov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblackbst

import (
	"math/rand"
	"testing"
)

type Int int

func (i Int) Compare(other KType) int {
	return int(i - other.(Int))
}

func TestCases(t *testing.T) {
	tree := New()
	tree.Put(Int(1), 1)
	tree.Put(Int(1), 1)
	if tree.Size() != 1 {
		t.Errorf("expecting len 1")
	}
	if !tree.Has(Int(1)) {
		t.Errorf("expecting to find key=1")
	}

	tree.Delete(Int(1))
	if tree.Size() != 0 {
		t.Errorf("expecting len 0")
	}
	if tree.Has(Int(1)) {
		t.Errorf("not expecting to find key=1")
	}

	tree.Delete(Int(1))
	if tree.Size() != 0 {
		t.Errorf("expecting len 0")
	}
	if tree.Has(Int(1)) {
		t.Errorf("not expecting to find key=1")
	}
}

func TestReverseInsertOrder(t *testing.T) {
	tree := New()
	n := 100
	for i := 0; i < n; i++ {
		tree.Put(Int(n-i), Int(n-i))
	}
	i := 0
	lo, _, _ := tree.Min()
	hi, _, _ := tree.Max()
	tree.RangedKeys(lo, hi, func(_ KType, v VType) bool {
		i++
		if v.(Int) != Int(i) {
			t.Errorf("bad order: got %d, expect %d", v.(Int), i)
		}
		return true
	})
}

func TestRange(t *testing.T) {
	tree := New()
	order := []K{
		"ab", "aba", "abc", "a", "aa", "aaa", "b", "a-", "a!",
	}
	for _, i := range order {
		tree.Put(i, i)
	}
	k := 0
	tree.RangedKeys(K("ab"), K("ac"), func(_ KType, v VType) bool {
		if k > 3 {
			t.Fatalf("returned more items than expected")
		}
		i1 := order[k]
		i2 := v.(K)
		if i1 != i2 {
			t.Errorf("expecting %s, got %s", i1, i2)
		}
		k++
		return true
	})
}

func TestRandomInsertOrder(t *testing.T) {
	tree := New()
	n := 1000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Put(Int(perm[i]), Int(perm[i]))
	}
	j := 0
	lo, _, _ := tree.Min()
	hi, _, _ := tree.Max()
	tree.RangedKeys(lo, hi, func(_ KType, v VType) bool {
		if v.(Int) != Int(j) {
			t.Fatalf("bad order")
		}
		j++
		return true
	})
}

func TestRandomReplace(t *testing.T) {
	tree := New()
	n := 100
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Put(Int(perm[i]), Int(perm[i]))
	}
	perm = rand.Perm(n)
	for i := 0; i < n; i++ {
		if replaced, _ := tree.Put(Int(perm[i]), Int(perm[i])); replaced == nil || replaced.(Int) != Int(perm[i]) {
			t.Errorf("error replacing")
		}
	}
}

func TestRandomInsertSequentialDelete(t *testing.T) {
	tree := New()
	n := 1000
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Put(Int(perm[i]), Int(perm[i]))
	}
	for i := 0; i < n; i++ {
		tree.Delete(Int(i))
	}
	if tree.Size() != 0 {
		printTreeStats(tree, "should be empty")
		t.Fatalf("tree has size %d", tree.Size())
	}
}

func TestRandomInsertDeleteNonExistent(t *testing.T) {
	tree := New()
	n := 100
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		tree.Put(Int(perm[i]), Int(perm[i]))
	}
	if _, ok := tree.Delete(Int(200)); ok {
		t.Errorf("deleted non-existent item")
	}
	if _, ok := tree.Delete(Int(-2)); ok {
		t.Errorf("deleted non-existent item")
	}
	for i := 0; i < n; i++ {
		if u, _ := tree.Delete(Int(i)); u == nil || u.(Int) != Int(i) {
			t.Errorf("delete failed")
		}
	}
	if _, ok := tree.Delete(Int(200)); ok {
		t.Errorf("deleted non-existent item")
	}
	if _, ok := tree.Delete(Int(-2)); ok {
		t.Errorf("deleted non-existent item")
	}
}

func TestRandomInsertPartialDeleteOrder(t *testing.T) {
	tree := New()
	n := 10
	perm := rand.Perm(n)
	for i := 0; i < n; i++ {
		vi := Int(perm[i])
		tree.Put(vi, vi)
	}
	for i := 1; i < n-1; i++ {
		vi := Int(i)
		_, ok := tree.Delete(vi)
		if !ok {
			t.Errorf("didn't delete %v", vi)
		}

	}
	j := 0
	lo, _, _ := tree.Min()
	hi, _, _ := tree.Max()
	tree.RangedKeys(lo, hi, func(_ KType, v VType) bool {
		t.Logf("v=%#v", v)
		switch j {
		case 0:
			if v.(Int) != Int(0) {
				t.Errorf("expecting 0")
			}
		case 1:
			if v.(Int) != Int(n-1) {
				t.Errorf("expecting %d, got %d", n-1, v)
			}
		}
		j++
		return true
	})
	if j != 2 {
		t.Errorf("should have ranged over only 2 keys, %d", j)
	}
}

func BenchmarkInsert(b *testing.B) {
	tree := New()
	for i := 0; i < b.N; i++ {
		tree.Put(Int(b.N-i), Int(b.N-i))
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	tree := New()
	for i := 0; i < b.N; i++ {
		tree.Put(Int(b.N-i), Int(b.N-i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(Int(i))
	}
}

func BenchmarkDeleteMin(b *testing.B) {
	b.StopTimer()
	tree := New()
	for i := 0; i < b.N; i++ {
		tree.Put(Int(b.N-i), Int(b.N-i))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.DeleteMin()
	}
}
