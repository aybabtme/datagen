# Command datagen

Generate datastructures for your types.

```bash
go get github.com/aybabtme/datagen/...
```

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

The following packages can be imported and used without code gen. However,
their location in the project is subject to change.

* `map/redblackbst` implements a red black balanced search tree,
based on the details provided in Algorithms 4th edition, by
Robert Sedgewick and Kevin Wayne. A red black bst is useful as
 a map that keeps its items in sorted order, while preserving
 efficient inserts, lookups and deletions.
* `set/redblackbst` is similar to the `map` implementation, but stores
no data about values.


## Contributions

See TODO file.

* Code should be `gofmt`'d.
* Code should have had `go vet` ran onto it.
* Code should have had `golint` ran onto it.

If you're not familiar with forks and contributions in Go, the flow
differs a bit from other languages.

While you can fork the project, you can't work under that fork's path
in your `GOPATH`. Instead, do the following.

* `go get github.com/aybabtme/datagen`
* `cd $GOPATH/src/github.com/aybabtme/datagen`
* Fork this project.
* `git remote add my_fork git@github.com:my_username/datagen.git`

If this seems odd to you, there's litterature elsewhere about why you
need to do this.

## Credits

The red black tree implementation is heavily inspired from the Java
implementation of Robert Sedgewick.

Some tests for the red black tree were extracted from GoLLRB, a similar
implementation by Petar Maymounkov.

