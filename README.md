# Command datagen

Generate datastructures for your types.

## With Go installed
```bash
$ go get -u github.com/aybabtme/datagen/...
```

## On linux

```bash
wget -qO- https://github.com/aybabtme/datagen/releases/download/0.1.4/datagen_Linux_x86_64.tar.gz | tar xvz
```

## On OS X

```bash
brew tap aybabtme/homebrew-tap
brew install datagen
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

## Supports

* Heap/Priority queues.
* Sorted maps.
* Sorted sets.
* Queues.

## Why

### Usability

Having datastructures that are specifically suited for your types is much easier
to code against than those relying on interfaces.

For instance, when you need a heap/priority queue:

```go
// this
func main() {
    h := NewIntHeap()
    for i := 20; i > 0; i-- {
        h.Push(i)
    }
}

// vs this (container/heap)
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    h := new(IntHeap)
    for i := 20; i > 0; i-- {
        h.Push(i)
    }
}
```

Or if you need a sorted map:

```go
// this
func main() {
    tree := NewSortedStringToStringMap()
    tree.Put("hello", "world")
    k, v, _ := tree.DeleteMin()
    log.Printf("%s, %s", k, v)
}

// vs this (GoLLRB)
type String struct {
    key string
    val string
}

func (s String) Less(than llrb.Item) bool {
    return string(s.key) < string(than.(String).key)
}

func main() {
    tree := llrb.New()
    tree.ReplaceOrInsert(String{key: "hello", val:"world"})
    kv := tree.DeleteMin().(String)
    log.Printf("%s, %s", kv.key, kv.val)
}
```


### Performance

In most case, a code generated datastructure will perform better than one that
uses interfaces.


#### SortedMap
                    | Operations  | GoLLRB ns/op | datagen ns/op | delta (smaller is better)
--------------------|-------------|--------------|---------------|---------------------------
`[]byte`:`string`   | Delete      | 3119         |  1487         | -52.32%
`float64`: `string` | Delete      | 2684         |  1797         | -33.05%
`int`:`string`      | Delete      | 2715         |  1236         | -54.48%
`string` `string`   | Delete      | 2830         |  1550         | -45.23%
`[]byte`:`string`   | DeleteMin   | 1002         |  1026         | +2.40%
`float64`: `string` | DeleteMin   | 1034         |  1256         | +21.47%
`int`:`string`      | DeleteMin   | 1045         |  1300         | +24.40%
`string` `string`   | DeleteMin   | 977          |  1231         | +26.00%
`[]byte`:`string`   | Insert      | 3606         |  1228         | -65.95%
`float64`: `string` | Insert      | 2722         |  1006         | -63.04%
`int`:`string`      | Insert      | 2736         |  1039         | -62.02%
`string` `string`   | Insert      | 3256         |  1121         | -65.57%

#### SortedSet

           | Operations   | GoLLRB ns/op | datagen ns/op | delta (smaller is better)
-----------|--------------|--------------|---------------|---------------------------
 `[]byte`  | Delete       | 2934         |  1620         | -44.79%
 `float64` | Delete       | 2712         |  1599         | -41.04%
 `int`     | Delete       | 2839         |  1003         | -64.67%
 `string`  | Delete       | 2811         |  1271         | -54.78%
 `[]byte`  | DeleteMin    | 999          |  1208         | +20.92%
 `float64` | DeleteMin    | 1049         |  1048         | -0.10%
 `int`     | DeleteMin    | 1000         |  1013         | +1.30%
 `string`  | DeleteMin    | 999          |  1014         | +1.50%
 `[]byte`  | Insert       | 3267         |  1214         | -62.84%
 `float64` | Insert       | 2705         |  899          | -66.77%
 `int`     | Insert       | 2720         |  869          | -68.05%
 `string`  | Insert       | 3152         |  1040         | -67.01%

#### Heap

             | Operations | stdlib ns/op   |  datagen ns/op   | delta  (smaller is better)
-------------|------------|----------------|------------------|---------------------------
`[]byte`     | Pop        | 1383           |  1323            | -4.34%
`float64`    | Pop        | 499            |  548             | +9.82%
`int`        | Pop        | 507            |  416             | -17.95%
`string`     | Pop        | 1081           |  1014            | -6.20%
`[]byte`     | Push       | 456            |  361             | -20.83%
`float64`    | Push       | 71.2           |  334             | +369.10%
`int`        | Push       | 70.2           |  255             | +263.25%
`string`     | Push       | 240            |  231             | -3.75%
`[]byte`     | various    | 9089651        |  8669574         | -4.62%
`float64`    | various    | 3697083        |  5099185         | +37.92%
`int`        | various    | 3686002        |  3685407         | -0.02%
`string`     | various    | 7612799        |  7007515         | -7.95%

#### Queue
             | Operations | stdlib ns/op   |  datagen ns/op   | delta  (smaller is better)
-------------|------------|----------------|------------------|---------------------------
`[]byte`     | Pop        | 72.1           | 60.3             | -16.37%
`float64`    | Pop        | 60.1           | 19.5             | -67.55%
`int`        | Pop        | 79.2           | 20.7             | -73.86%
`string`     | Pop        | 79.6           | 51.3             | -35.55%
`[]byte`     | Push       | 326            | 175              | -46.32%
`float64`    | Push       | 141            | 18.9             | -86.60%
`int`        | Push       | 218            | 21.7             | -90.05%
`string`     | Push       | 201            | 56.8             | -71.74%
`[]byte`     | Serial     | 395            | 229              | -42.03%
`float64`    | Serial     | 197            | 37.5             | -80.96%
`int`        | Serial     | 268            | 37.0             | -86.19%
`string`     | Serial     | 215            | 165              | -23.26%
`[]byte`     | TickTock   | 217            | 38.9             | -82.07%
`float64`    | TickTock   | 128            | 32.3             | -74.77%
`int`        | TickTock   | 152            | 32.1             | -78.88%
`string`     | TickTock   | 165            | 36.4             | -77.94%


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
* `heap` is a heap implementation inspired from Algorithms 4th edition and
the `container/heap` implementation.
* `queue` is a queue implementation adapted from github.com/eapachae/queue.

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

The red black tree and heap implementations are heavily inspired from the Java
implementations of Robert Sedgewick.

The heap implementation was inspired, and the comments/tests adapted from `container/heap`.

The queue implementation was adapted from `github.com/eapache/queue`, a package by
Evan Huus.

Some tests for the red black tree were extracted from GoLLRB, a similar
implementation by Petar Maymounkov.
