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

func sortedMap() cli.Command {

	kTypeFlag := cli.StringFlag{
		Name:  "key",
		Usage: "type that will be used for keys",
	}
	vTypeFlag := cli.StringFlag{
		Name:  "val",
		Usage: "type that will be used for values",
	}

	return cli.Command{
		Name:      "sorted-map",
		ShortName: "smap",
		Usage:     "Create a sorted map customized for your types.",
		Description: `Create a sorted map customized for your types. The map is built
on a left leaning red black balanced search tree. The implementation has good
performance and is well tested, with 100% test coverage. (the tests are not
generated with the custom type)`,
		Flags: []cli.Flag{kTypeFlag, vTypeFlag},
		Action: func(ctx *cli.Context) {
			ktype := valOrDefault(ctx, kTypeFlag)
			vtype := valOrDefault(ctx, vTypeFlag)

			kname := ktype
			vname := vtype
			if kname == "[]byte" {
				kname = "Bytes"
			}
			if []byte(kname)[0] == '*' {
				kname = kname[1:]
			}
			if vname == "[]byte" {
				vname = "Bytes"
			}
			if []byte(vname)[0] == '*' {
				vname = vname[1:]
			}
			typeName := fmt.Sprintf("Sorted%sTo%sMap", strings.Title(kname), strings.Title(vname))
			nodeName := fmt.Sprintf("node%sTo%s", strings.Title(kname), strings.Title(vname))

			cwd, _ := os.Getwd()
			pkgname := fmt.Sprintf("package %s", filepath.Base(cwd))

			src := []byte(redblackbstMap)
			src = bytes.Replace(src, []byte("package redblackbst"), []byte(pkgname), 1)

			// need to replace Compare before replacing KType
			src = replaceCompareFunc(ktype, src)
			src = bytes.Replace(src, []byte("KType"), []byte(ktype), -1)
			src = bytes.Replace(src, []byte("VType"), []byte(vtype), -1)
			src = bytes.Replace(src, []byte("RedBlack"), []byte(typeName), -1)
			src = bytes.Replace(src, []byte("mapnode"), []byte(nodeName), -1)

			fmt.Println(string(src))
		},
	}
}

func replaceCompareFunc(ktype string, src []byte) []byte {
	var tmpl string
	orig := "func (r RedBlack) compare(a, b KType) int { return a.Compare(b) }"

	switch ktype {

	default:
		// don't change anything by default
		return src

	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		tmpl = "func (r RedBlack) compare(a, b KType) int { return a-b }"

	case "float32", "float64":
		tmpl = `
const e = 0.00000001

func (r RedBlack) compare(a, b KType) int {
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
func (r RedBlack) compare(a, b KType) int {
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
func  (r RedBlack) compare(a, b KType) int { return bytes.Compare(a, b) }
`

	}
	return bytes.Replace(src, []byte(orig), []byte(tmpl), -1)
}
