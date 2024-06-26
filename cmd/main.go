package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	exact "github.com/williammartin/gh-exact"
	"github.com/williammartin/gh-exact/cli"
	"github.com/williammartin/gh-exact/fs"
)

func main() {
	var filePath string

	cmdSave := &cobra.Command{
		Use:   "dump",
		Short: "Dump the extensions to a file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			app := exact.App{
				ExtensionManager: cli.ExtensionManager{
					Out: os.Stdout,
					Err: os.Stderr,
				},
				ManifestStorage: fs.ManifestStorage{
					FilePath: filePath,
				},
			}

			return app.Dump()
		},
	}

	var pin bool
	cmdRestore := &cobra.Command{
		Use:   "restore",
		Short: "Restore extensions from a file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			app := exact.App{
				ExtensionManager: cli.ExtensionManager{
					Out: os.Stdout,
					Err: os.Stderr,
				},
				ManifestStorage: fs.ManifestStorage{
					FilePath: filePath,
				},
			}

			versioning := exact.LatestVersions
			if pin {
				versioning = exact.PinVersions
			}

			return app.Restore(versioning)
		},
	}
	cmdRestore.Flags().BoolVarP(&pin, "pin", "p", false, "Pin the extension versions")

	rootCmd := &cobra.Command{Use: "gh-exact"}
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "extensions.yaml", "extension file path (default: extensions.yaml)")

	rootCmd.AddCommand(cmdSave, cmdRestore)
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
