To benchmark implementations, use tags:

* `-tags=own`: our code generated implementations.
* `-tags=other`: known good interface{} based libraries.

The prefix (`#` in `#_name_test.go`) are to keep the order of execution of the
benchmarks the same, to make it easy to work with `benchcmp`.
