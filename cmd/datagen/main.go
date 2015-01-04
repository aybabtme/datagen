package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("datagen: ")
	app := cli.NewApp()

	app.Name = "datagen"
	app.Email = "antoinegrondin@gmail.com"
	app.Author = "Antoine Grondin"
	app.Version = "0.1"
	app.Commands = append(app.Commands, sortedMap())

	app.Run(os.Args)
}

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
			typeName := fmt.Sprintf("Sorted%sTo%sMap", strings.Title(ktype), strings.Title(vtype))

			cwd, _ := os.Getwd()
			pkgname := fmt.Sprintf("package %s", filepath.Base(cwd))

			src := []byte(redblackbst)
			src = bytes.Replace(src, []byte("package redblackbst"), []byte(pkgname), 1)
			src = bytes.Replace(src, []byte("KType"), []byte(ktype), -1)
			src = bytes.Replace(src, []byte("VType"), []byte(vtype), -1)
			src = bytes.Replace(src, []byte("RedBlack"), []byte(typeName), -1)

			switch ktype {
			case "int", "int8", "int16", "int32", "int64",
				"uint", "uint8", "uint16", "uint32", "uint64",
				"float32", "float64":
				src = bytes.Replace(src, []byte(".Compare(h.key)"), []byte("-h.key"), -1)
			}

			fmt.Println(string(src))
		},
	}
}

func valOrDefault(ctx *cli.Context, f cli.StringFlag) string {
	str := ctx.String(f.Name)
	if str != "" {
		return str
	}
	if f.Value == "" {
		log.Printf("flag not set: %q", f.Name)
		cli.ShowAppHelp(ctx)
		os.Exit(1)
	}
	return f.Value
}
