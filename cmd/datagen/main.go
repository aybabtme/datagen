// Command datagen
//
// Generate datastructures for your types.
//
// Builds upon well tested implementations of datastructures
// to generate customized implementations for your types.
// Alike to what you would get with generics, but with code
// generation instead.
//
// You can use it manually or with `go generate`.
//
// For more information, invoke the command with the `-h` flag.
//
// Alike to `go generate` and other code gen tools, this tool
// is meant for package authors who wish to generate code.
// It should not be used as a build step for your users.
package main

import (
	"log"
	"os"

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
	app.Commands = append(app.Commands, sortedSet())

	app.Run(os.Args)
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
