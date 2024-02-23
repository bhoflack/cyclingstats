package main

import (
	"github.com/spf13/cobra"
)

func cyclingstatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cyclingstats",
		Short: "cyclingstats",
	}

	cmd.AddCommand(fetchCommand())
	cmd.AddCommand(cyclistCmd())
	cmd.AddCommand(statCmd())

	return cmd
}

func main() {
	cmd := cyclingstatsCmd()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
