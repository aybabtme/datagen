package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

func queue() cli.Command {

	keyTypeFlag := cli.StringFlag{
		Name:  "key",
		Usage: "type that will be held in the queue",
	}

	return cli.Command{
		Name:      "queue",
		ShortName: "q",
		Usage:     "Create a queue (list) customized for your types.",
		Description: `Create a queue customized for your types. The implementation
is based on a ring buffer, which has good performance and is well tested.
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

			typeName := fmt.Sprintf("%sQueue", strings.Title(kname))

			cwd, _ := os.Getwd()
			pkgname := fmt.Sprintf("package %s", filepath.Base(cwd))

			src := []byte(queueSrc)
			src = bytes.Replace(src, []byte("package queue"), []byte(pkgname), 1)

			src = bytes.Replace(src, []byte("nilKType"), []byte("nil"+kname), -1) // before KType's replace
			src = bytes.Replace(src, []byte("KType"), []byte(ktype), -1)
			src = bytes.Replace(src, []byte("Queue"), []byte(typeName), -1)

			fmt.Println(string(src))
		},
	}
}
