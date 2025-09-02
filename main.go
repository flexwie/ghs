package main

import (
	"context"
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/cli/go-gh/v2"
	"github.com/flexwie/ghs/pkg"
	"github.com/spf13/cobra"
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
	root := &cobra.Command{
		Use:                   "ghs [gist] [args]",
		Short:                 "npx-like script execution for GitHub gists",
		DisableFlagParsing:    true,
		SilenceUsage:          true,
		SilenceErrors:         true,
		Args:                  cobra.ArbitraryArgs,
		DisableFlagsInUseLine: true,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			log.Debug(toComplete)

			outputs := []string{"hello", "moto"}
			return outputs, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := gh.Path()
			if err != nil {
				return errors.New("gh cli is not installed")
			}

			if len(args) == 0 {
				return errors.New("gist is required")
			}

			gist := args[0]
			if len(gist) == 0 {
				return errors.New("gist is required")
			}

			gistArgs := args[1:]
			log.Debug("parsing args", "gist", gist, "args", gistArgs)

			return pkg.ExecGist(context.Background(), gist, gistArgs)
		},
	}

	completion := &cobra.Command{
		Use:                   "completion [shell]",
		Short:                 "Generate autocompletion script",
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				_ = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				_ = cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			case "fish":
				_ = cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			case "powershell":
				_ = cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
			}

			return nil
		},
	}

	root.AddCommand(completion)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
