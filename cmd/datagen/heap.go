package main

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func heap() cli.Command {

	keyTypeFlag := cli.StringFlag{
		Name:  "key",
		Usage: "type that will be held in the heap",
	}

	return cli.Command{
		Name:      "heap",
		ShortName: "pq",
		Usage:     "Create a heap (priority queue) customized for your types.",
		Description: `Create a heap customized for your types. The implementation
has good performance and is well tested, with 100% test coverage.
(the tests are not generated with the custom type)`,
		Flags: []cli.Flag{keyTypeFlag},
		Action: func(ctx *cli.Context) {
			ktype := valOrDefault(ctx, keyTypeFlag)

			kname := ktype
			if len(kname) > 1 && []byte(kname)[0] == '*' {
				kname = kname[1:]
			}
			if len(kname) > 2 && kname[:2] == "[]" {
				kname = strings.Title(kname[2:]) + "s"
			}

			typeName := fmt.Sprintf("%sHeap", strings.Title(kname))

			cwd, _ := os.Getwd()
			pkgname := fmt.Sprintf("package %s", filepath.Base(cwd))

			src := []byte(heapSrc)
			src = bytes.Replace(src, []byte("package heap"), []byte(pkgname), 1)

			// need to replace Compare before replacing KType
			src = replaceHeapCompareFunc(ktype, src)
			src = bytes.Replace(src, []byte("KType"), []byte(ktype), -1)
			src = bytes.Replace(src, []byte("Heap"), []byte(typeName), -1)

			fmt.Println(string(src))
		},
	}
}

func replaceHeapCompareFunc(ktype string, src []byte) []byte {
	var tmpl string
	orig := "func (h Heap) compare(a, b KType) int { return a.Compare(b) }"

	switch ktype {

	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		tmpl = "func (h Heap) compare(a, b KType) int { return int(a) - int(b) }"

	case "float32", "float64":
		tmpl = `
const e = 0.00000001

func (h Heap) compare(a, b KType) int {
    diff := (a-b)/a
    if diff < -e {
        return -1
    } else if diff > e {
        return 1
    }
    return 0
}`

	case "string":
		tmpl = `
func (h Heap) compare(a, b KType) int {
    if a < b {
        return -1
    }
    if a > b {
        return 1
    }
    return 0
}`

	case "[]byte":
		log.Printf("WARNING: using []byte as keys can lead to undefined behavior if the []byte are modified after insertion!!!")
		tmpl = `
// WARNING: using []byte as keys can lead to undefined behavior if the
// []byte are modified after insertion!!!
func  (h Heap) compare(a, b KType) int { return bytes.Compare(a, b) }
`

	default:

		// if storing slices, use `len()` for comparison
		if len(ktype) > 2 && ktype[:2] == "[]" {
			log.Printf("%s: order will be determined based on value of len(%s)", ktype, ktype)
			tmpl = fmt.Sprintf(
				"func (h Heap) compare(a, b %s) int { return len(a)-len(b) }",
				ktype,
			)
		} else {
			// otherwise don't change anything by default, let the user
			// provide a `Compare` func
			log.Printf("type %q will need to implement a Compare func: %s",
				ktype,
				fmt.Sprintf(`
	func (%[1]s %s) Compare(other %s) int {
		if %[1]s > other {
			return 1
		} else if %[1]s < other {
			return -1
		}
		return 0
	}`, strings.ToLower(ktype[:1]), ktype, ktype))
			return src
		}

	}

	return bytes.Replace(src, []byte(orig), []byte(tmpl), -1)
}
