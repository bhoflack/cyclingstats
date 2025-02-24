package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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
	cmd.AddCommand(removeCyclistCmd())
	cmd.AddCommand(addCyclisInteractiveCmd())

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

func addCyclisInteractiveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-interactive",
		Short: "add-interactive",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}
			defer client.Close()

			current, err := client.ListCyclists()
			if err != nil {
				return err
			}

			fmt.Printf("Current cyclists:\n")
			for _, cyclist := range current {
				fmt.Printf("%v\n", cyclist)
			}

			for {
				// ask for the name
				fmt.Printf("Name:\n")

				var name string
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					name = scanner.Text()
				}

				if name == "" {
					break
				}

				cyclists, err := procyclingstats.Find(name)
				if err != nil {
					return err
				}

				if len(cyclists) == 0 {
					log.Printf("no cyclist found for %s", name)
					continue
				}

				if len(cyclists) == 1 {
					if err := client.AddCyclist(cyclists[0]); err != nil {
						return err
					}
					continue
				}

				for i, cyclist := range cyclists {
					fmt.Printf("%d: %v\n", i, cyclist)
				}

				fmt.Printf("Select the cyclist:\n")
				n := -1
				if _, err := fmt.Scanln(&n); err != nil {
					return nil
				}

				if n < 0 || n >= len(cyclists) {
					log.Printf("invalid selection")
					continue
				}

				if err := client.AddCyclist(cyclists[n]); err != nil {
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

func removeCyclistCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}
			defer client.Close()

			for _, name := range args {
				if err := client.RemoveCyclist(name); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return cmd
}
