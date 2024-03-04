package cmd

import (
	"github.com/spf13/cobra"
	"paranoidbackup/go-paper/internal/di"
)

var (
	path    string
	project string
	workDir string

	rootCmd = &cobra.Command{
		Use:   "go-paper",
		Short: "go-paper",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := di.Bootstrap(workDir)
			if err != nil {
				panic(err)
			}
			defer app.HandlePanic()

			task := app.BackupTask()

			if project == "" {
				err = task.BackupWithNewProject(path)
				if err != nil {
					panic(err)
				}
			} else {
				err = task.BackupWithExistingProject(project, path)
				if err != nil {
					panic(err)
				}
			}

			app.Halt()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&path, "path", "p", "", "Path to file for backup")
	rootCmd.Flags().StringVarP(&project, "project", "i", "", "Project ID to reuse")
	rootCmd.Flags().StringVarP(&workDir, "work-dir", "d", "", "Working directory where to get/save project files and save backup")
}
