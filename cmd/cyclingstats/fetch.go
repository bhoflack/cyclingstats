package main

import (
	"github.com/bhoflack/cyclingstats/pkg/db"
	"github.com/bhoflack/cyclingstats/pkg/procyclingstats"
	"github.com/spf13/cobra"
)

func fetchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "fetch",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}
			cyclists, err := client.ListCyclists()
			if err != nil {
				return err
			}

			if err := procyclingstats.UpdateStats(cyclists); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
