// Package suite benchmarks datagen's datastructures against their generic
// ports and against third party implementations.
//
// Every implementation runs the same benchmark body: a structure holds a
// fixed number of elements (see sizes) while it is measured, so that faster
// implementations aren't penalized with larger structures. Each
// implementation plugs in through a small adapter, see the impl_*_test.go
// files.
//
// Benchmarks are named key=<type>/n=<size>/op=<operation>/impl=<name>, so
// implementations can be compared with:
//
//	go test -run '^$' -bench . -count 6 > out.bench
//	benchstat -col /impl out.bench
package suite
