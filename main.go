package main

import (
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/charmbracelet/log"
	"github.com/flexwie/ghs/pkg"
	"github.com/urfave/cli/v2"
)

type config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"warn"`
}

var start time.Time = time.Now()

func init() {
	c := config{}
	err := env.ParseWithOptions(&c, env.Options{
		Prefix: "GHS_",
	})
	if err != nil {
		panic(err)
	}

	level, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		panic(err)
	}

	var logger *log.Logger = log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: false,
		ReportCaller:    false,
		Level:           level,
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
			defer log.Debug("finished execution of script", "duration", time.Since(start))
			return pkg.Run(ctx.Args().Slice())
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
