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
		UseShortOptionHandling: true,
		Flags: []cli.Flag{&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Value:   true,
			Usage:   "only print output of your command",
		}},
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

			return pkg.ExecGist(context.Background(), gist)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
