package cmd

import (
	"go-paper/go-paper/internal/app"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-paper",
		Short: "go-paper",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := app.Bootstrap()
			if err != nil {
				panic(err)
			}
			defer app.HandlePanic()

			app.Halt()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
