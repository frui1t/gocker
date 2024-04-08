package main

import (
	"fmt"
	"gocker/gocker/version"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	runGocker := newGockerCommand()
	if err := runGocker.Execute(); err != nil {
		fmt.Printf("Gocker init fail,err: %v", err)
		os.Exit(1)
	}
}
func newGockerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "Gocker [OPTIONS] COMMAND [ARG...]",
		Short:            "A self-sufficient runtime for containers",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("Gocker: '%s' is not a Gocker command.\nSee 'Gocker --help'", args[0])
		},
		Version:               fmt.Sprintf("%s, build %s, built %s", version.Version, version.GitCommit, version.BuildTime),
		DisableFlagsInUseLine: true,
	}
	return cmd
}
