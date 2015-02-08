package queue

import "testing"

func TestQueueLen(t *testing.T) {
	q := NewQueue(0)

	if q.Len() != 0 {
		t.Error("empty queue length not 0")
	}

	for i := 0; i < 1000; i++ {
		q.Push(i)
		if q.Len() != i+1 {
			t.Error("adding: queue with", i, "elements has length", q.Len())
		}
	}
	for i := 0; i < 1000; i++ {
		q.Pop()
		if q.Len() != 1000-i-1 {
			t.Error("removing: queue with", 1000-i-i, "elements has length", q.Len())
		}
	}
}

func TestQueueGet(t *testing.T) {
	q := NewQueue(0)

	for i := 0; i < 1000; i++ {
		q.Push(i)
		for j := 0; j < q.Len(); j++ {
			if q.Get(j).(int) != j {
				t.Errorf("index %d doesn't contain %d", j, j)
			}
		}
	}
}

func TestQueueGetOutOfRangePanics(t *testing.T) {
	q := NewQueue(0)

	q.Push(1)
	q.Push(2)
	q.Push(3)

	assertPanics(t, "should panic when negative index", func() {
		q.Get(-1)
	})

	assertPanics(t, "should panic when index greater than length", func() {
		q.Get(4)
	})
}

func TestQueuePeekOutOfRangePanics(t *testing.T) {
	q := NewQueue(0)

	assertPanics(t, "should panic when peeking empty queue", func() {
		q.Peek()
	})

	q.Push(1)
	q.Peek() // should not panic
	q.Pop()

	assertPanics(t, "should panic when peeking emptied queue", func() {
		q.Peek()
	})
}

func TestQueuePopOutOfRangePanics(t *testing.T) {
	q := NewQueue(0)

	assertPanics(t, "should panic when removing empty queue", func() {
		q.Pop()
	})

	q.Push(1)
	q.Pop()

	assertPanics(t, "should panic when removing emptied queue", func() {
		q.Pop()
	})
}

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		} else {
			t.Logf("%s: got panic as expected: %v", name, r)
		}
	}()

	f()
}

// General warning: Go's benchmark utility (go test -bench .) increases the number of
// iterations until the benchmarks take a reasonable amount of time to run; memory usage
// is *NOT* considered. On my machine, these benchmarks hit around ~1GB before they've had
// enough, but if you have less than that available and start swapping, then all bets are off.

func BenchmarkQueueSerial(b *testing.B) {
	q := NewQueue(0)
	for i := 0; i < b.N; i++ {
		q.Push(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkQueueGet(b *testing.B) {
	q := NewQueue(0)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Get(i)
	}
}

func BenchmarkQueueTickTock(b *testing.B) {
	q := NewQueue(0)
	for i := 0; i < b.N; i++ {
		q.Push(nil)
		q.Pop()
	}
}
