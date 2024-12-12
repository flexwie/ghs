package main

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
	"github.com/flexwie/ghs/pkg"
	"github.com/urfave/cli/v2"
)

func init() {
	var logger *log.Logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: false,
		ReportCaller:    false,
		Level:           log.InfoLevel,
	})

	log.SetDefault(logger)
}

func main() {
	app := &cli.App{
		Name:                   "ghs",
		Usage:                  "npx-like script execution for GitHub gists",
		SkipFlagParsing:        true,
		UseShortOptionHandling: true,
		UsageText:              "ghs <gist name> [arguments/flags...]",
		Authors: []*cli.Author{
			{Name: "Felix Wieland", Email: "ghs@felixwie.com"},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Bool("quiet") {
				log.SetLevel(log.ErrorLevel)
			}

			gist := ctx.Args().Get(0)
			if len(gist) == 0 {
				return cli.Exit("no gist provided", 2)
			}

			args := ctx.Args().Slice()
			return pkg.ExecGist(context.Background(), gist, args[1:])
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
