# Command datagen

Generate datastructures for your types.

Builds upon well tested implementations of datastructures
to generate customized implementations for your types.
Alike to what you would get with generics, but with code
generation instead.

You can use it manually or with `go generate`.

For more information, invoke the command with the `-h` flag.

Alike to `go generate` and other code gen tools, this tool
is meant for package authors who wish to generate code.
It should not be used as a build step for your users.

## Subpackages

* `redblackbst` implements a red black balanced search tree,
based on the details provided in Algorithms 4th edition, by
Robert Sedgewick and Kevin Wayne. A red black bst is useful as
 a map that keeps its items in sorted order, while preserving
 efficient inserts, lookups and deletions.

## Credits

The red black tree implementation is heavily inspired from the Java
implementation of Robert Sedgewick.

Some tests for the red black tree were extracted from GoLLRB, a similar
implementation by Petar Maymounkov.
