package main

import (
	"bytes"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
	"strings"
)

func sortedSet() cli.Command {

	keyTypeFlag := cli.StringFlag{
		Name:  "key",
		Usage: "type that will be held in the set",
	}

	return cli.Command{
		Name:      "sorted-set",
		ShortName: "sset",
		Usage:     "Create a sorted set customized for your types.",
		Description: `Create a sorted set customized for your types. The set is built
on a left leaning red black balanced search tree. The implementation has good
performance and is well tested, with 100% test coverage. (the tests are not
generated with the custom type)`,
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

			typeName := fmt.Sprintf("Sorted%sSet", strings.Title(kname))
			nodeName := fmt.Sprintf("node%s", strings.Title(kname))

			cwd, _ := os.Getwd()
			pkgname := fmt.Sprintf("package %s", filepath.Base(cwd))

			src := []byte(redblackbstSetSrc)
			src = bytes.Replace(src, []byte("package redblackbst"), []byte(pkgname), 1)

			// need to replace Compare before replacing KType
			src = replaceRbstCompareFunc(ktype, src)
			src = bytes.Replace(src, []byte("KType"), []byte(ktype), -1)
			src = bytes.Replace(src, []byte("RedBlack"), []byte(typeName), -1)
			src = bytes.Replace(src, []byte("treenode"), []byte(nodeName), -1)

			fmt.Println(string(src))
		},
	}
}
