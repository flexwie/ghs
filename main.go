package main

import (
	"context"
	"log"
	"os"

	"github.com/flexwie/ghs/pkg"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "ghs",
		Usage:                  "npx-like script execution for GitHub gists",
		UseShortOptionHandling: true,
		Authors: []*cli.Author{
			{Name: "Felix Wieland", Email: "ghs@felixwie.com"},
		},
		Action: func(ctx *cli.Context) error {
			gist := ctx.Args().Get(0)
			if len(gist) == 0 {
				return cli.Exit("no gist provided", 2)
			}

			return pkg.ExecGist(context.Background(), gist)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
