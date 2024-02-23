package main

import (
	"log"

	"github.com/bhoflack/cyclingstats/pkg/db"
	"github.com/bhoflack/cyclingstats/pkg/procyclingstats"
	"github.com/spf13/cobra"
)

func cyclistCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cyclist",
		Short: "cyclist",
	}

	cmd.AddCommand(addCyclistCmd())
	cmd.AddCommand(listCyclistsCmd())

	return cmd
}

func addCyclistCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}
			defer client.Close()

			for _, name := range args {
				cyclists, err := procyclingstats.Find(name)
				if err != nil {
					return err
				}

				if len(cyclists) == 0 {
					log.Printf("no cyclist found for %s", name)
					continue
				}

				if len(cyclists) != 1 {
					log.Printf("found %v for %s,  will pick %v", cyclists, name, cyclists[0])
				}

				if err := client.AddCyclist(cyclists[0]); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return cmd
}

func listCyclistsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}
			defer client.Close()

			cyclists, err := client.ListCyclists()
			if err != nil {
				return err
			}

			for _, cyclist := range cyclists {
				cmd.Println(cyclist)
			}

			return nil
		},
	}

	return cmd
}
