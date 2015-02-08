// +build other

package bench

import (
	"container/list"
	"testing"
)

// String

func Benchmark_Queue_String_Serial(b *testing.B) {
	vals := makeStrings(b.N)
	q := list.New()
	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_String_TickTock(b *testing.B) {
	vals := makeStrings(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_String_Push(b *testing.B) {
	vals := makeStrings(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
}

func Benchmark_Queue_String_Pop(b *testing.B) {
	vals := makeStrings(b.N)
	q := list.New()
	for _, v := range vals {
		q.PushBack(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

// Int

func Benchmark_Queue_Int_Serial(b *testing.B) {
	vals := makeInts(b.N)
	q := list.New()
	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_Int_TickTock(b *testing.B) {
	vals := makeInts(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_Int_Push(b *testing.B) {
	vals := makeInts(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
}

func Benchmark_Queue_Int_Pop(b *testing.B) {
	vals := makeInts(b.N)
	q := list.New()
	for _, v := range vals {
		q.PushBack(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

// Float64

func Benchmark_Queue_Float64_Serial(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := list.New()
	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_Float64_TickTock(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_Float64_Push(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
}

func Benchmark_Queue_Float64_Pop(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := list.New()
	for _, v := range vals {
		q.PushBack(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

// Bytes

func Benchmark_Queue_Bytes_Serial(b *testing.B) {
	vals := makeBytess(b.N)
	q := list.New()
	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_Bytes_TickTock(b *testing.B) {
	vals := makeBytess(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
		q.Remove(q.Front())
	}
}

func Benchmark_Queue_Bytes_Push(b *testing.B) {
	vals := makeBytess(b.N)
	q := list.New()

	b.ResetTimer()
	for _, v := range vals {
		q.PushBack(v)
	}
}

func Benchmark_Queue_Bytes_Pop(b *testing.B) {
	vals := makeBytess(b.N)
	q := list.New()
	for _, v := range vals {
		q.PushBack(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Remove(q.Front())
	}
}
