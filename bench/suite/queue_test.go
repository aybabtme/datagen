package suite

import "testing"

type fifo[K key] interface {
	Push(k K)
	Pop() K
}

// queueOps:
//   - filldrain: push n keys then pop them all, reusing the same queue
//     across iterations (sec/op is per round, ns/key per push+pop).
//   - ticktock: with n keys queued, push one and pop one.
var queueOps = []string{"filldrain", "ticktock"}

func queueImpl[Q fifo[K], K key](name string, newQueue func() Q) impl[K] {
	return impl[K]{
		name:  name,
		run:   func(b *testing.B, op string, keys []K) { benchQueue(b, op, newQueue, keys) },
		check: func(t *testing.T, keys []K) { checkQueue(t, newQueue, keys) },
	}
}

func benchQueue[Q fifo[K], K key](b *testing.B, op string, newQueue func() Q, keys []K) {
	q := newQueue()
	switch op {
	case "filldrain":
		for b.Loop() {
			for _, k := range keys {
				q.Push(k)
			}
			for range keys {
				q.Pop()
			}
		}
		perKey(b, len(keys))
	case "ticktock":
		for _, k := range keys {
			q.Push(k)
		}
		i := 0
		for b.Loop() {
			q.Push(keys[i])
			q.Pop()
			i = next(i, len(keys))
		}
	default:
		b.Fatalf("unknown op %q", op)
	}
}

var (
	intQueues = []impl[int]{
		queueImpl("own", newOwnIntQueue),
		queueImpl("generic", newGenericQueue[int]),
		queueImpl("stdlib", newStdlibQueue[int]),
		queueImpl("gammazero", newGammazeroQueue[int]),
		queueImpl("eapache", newEapacheQueue[int]),
	}
	stringQueues = []impl[string]{
		queueImpl("own", newOwnStringQueue),
		queueImpl("generic", newGenericQueue[string]),
		queueImpl("stdlib", newStdlibQueue[string]),
		queueImpl("gammazero", newGammazeroQueue[string]),
		queueImpl("eapache", newEapacheQueue[string]),
	}
)

func BenchmarkQueue(b *testing.B) { suite(b, queueOps, intQueues, stringQueues) }
