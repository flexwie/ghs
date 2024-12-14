package main

import (
	"context"
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/cli/go-gh/v2"
	"github.com/flexwie/ghs/pkg"
	"github.com/urfave/cli/v2"
)

func init() {
	var logger *log.Logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: false,
		ReportCaller:    false,
		Level:           log.DebugLevel,
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
			_, err := gh.Path()
			if err != nil {
				return errors.New("gh cli is not installed")
			}

			gist := ctx.Args().Get(0)
			if len(gist) == 0 {
				return cli.Exit("no gist provided", 2)
			}

			args := ctx.Args().Slice()[1:]
			log.Debug("parsing args", "gist", gist, "args", args)

			return pkg.ExecGist(context.Background(), gist, args)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
