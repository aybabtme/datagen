// +build own

package bench

import (
	"testing"

	. "github.com/aybabtme/datagen/codegen"
)

// String

func Benchmark_Queue_String_Serial(b *testing.B) {
	vals := makeStrings(b.N)
	q := NewStringQueue(0)
	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func Benchmark_Queue_String_TickTock(b *testing.B) {
	vals := makeStrings(b.N)
	q := NewStringQueue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
		q.Pop()
	}
}

func Benchmark_Queue_String_Push(b *testing.B) {
	vals := makeStrings(b.N)
	q := NewStringQueue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
}

func Benchmark_Queue_String_Pop(b *testing.B) {
	vals := makeStrings(b.N)
	q := NewStringQueue(0)
	for _, v := range vals {
		q.Push(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

// Int

func Benchmark_Queue_Int_Serial(b *testing.B) {
	vals := makeInts(b.N)
	q := NewIntQueue(0)
	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func Benchmark_Queue_Int_TickTock(b *testing.B) {
	vals := makeInts(b.N)
	q := NewIntQueue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
		q.Pop()
	}
}

func Benchmark_Queue_Int_Push(b *testing.B) {
	vals := makeInts(b.N)
	q := NewIntQueue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
}

func Benchmark_Queue_Int_Pop(b *testing.B) {
	vals := makeInts(b.N)
	q := NewIntQueue(0)
	for _, v := range vals {
		q.Push(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

// Float64

func Benchmark_Queue_Float64_Serial(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := NewFloat64Queue(0)
	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func Benchmark_Queue_Float64_TickTock(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := NewFloat64Queue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
		q.Pop()
	}
}

func Benchmark_Queue_Float64_Push(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := NewFloat64Queue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
}

func Benchmark_Queue_Float64_Pop(b *testing.B) {
	vals := makeFloat64s(b.N)
	q := NewFloat64Queue(0)
	for _, v := range vals {
		q.Push(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

// Bytes

func Benchmark_Queue_Bytes_Serial(b *testing.B) {
	vals := makeBytess(b.N)
	q := NewBytesQueue(0)
	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func Benchmark_Queue_Bytes_TickTock(b *testing.B) {
	vals := makeBytess(b.N)
	q := NewBytesQueue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
		q.Pop()
	}
}

func Benchmark_Queue_Bytes_Push(b *testing.B) {
	vals := makeBytess(b.N)
	q := NewBytesQueue(0)

	b.ResetTimer()
	for _, v := range vals {
		q.Push(v)
	}
}

func Benchmark_Queue_Bytes_Pop(b *testing.B) {
	vals := makeBytess(b.N)
	q := NewBytesQueue(0)
	for _, v := range vals {
		q.Push(v)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
