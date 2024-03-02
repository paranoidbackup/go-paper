package cmd

import (
	"github.com/spf13/cobra"
	"paranoidbackup/go-paper/internal/di"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-paper",
		Short: "go-paper",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := di.Bootstrap("")
			if err != nil {
				panic(err)
			}
			defer app.HandlePanic()

			app.BackupTask() // TODO: call method here

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
