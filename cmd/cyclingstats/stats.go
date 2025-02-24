package main

import (
	"strings"

	"github.com/bhoflack/cyclingstats/pkg/db"
	"github.com/bhoflack/cyclingstats/pkg/procyclingstats"
	"github.com/spf13/cobra"
)

func statCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stat",
		Short: "stat",
	}

	cmd.AddCommand(listStatsCmd())
	cmd.AddCommand(upcomingRacesCmd())
	cmd.AddCommand(participatingCmd())

	return cmd
}

func upcomingRacesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upcoming",
		Short: "upcoming",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}

			pages, err := client.GetPages()
			if err != nil {
				return err
			}

			for rider, page := range pages {
				stats, err := procyclingstats.Extract(page)
				if err != nil {
					return err
				}

				cmd.Printf("%s: %v\n", rider, stats.UpcomingRaces)
			}
			return nil
		},
	}
	return cmd
}

func participatingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "participating",
		Short: "participating",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()

			pages, err := client.GetPages()
			if err != nil {
				return err
			}

			ridersToRaces := make(map[string][]string)
			for rider, page := range pages {
				stats, err := procyclingstats.Extract(page)
				if err != nil {
					return err
				}

				ridersToRaces[rider] = stats.UpcomingRaces
			}

			racesWithRiders := make(map[string][]string)
			for _, race := range args {
				for rider, races := range ridersToRaces {
					for _, r := range races {
						if strings.HasPrefix(r, race) {
							racesWithRiders[race] = append(racesWithRiders[race], rider)
						}
					}
				}
			}

			for race, riders := range racesWithRiders {

				cmd.Printf("%s: %v\n", race, riders)
			}

			return nil
		},
	}

	return cmd
}

func listStatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.NewClient()
			if err != nil {
				return err
			}

			stats, err := client.GetPages()
			if err != nil {
				return err
			}

			for _, stat := range stats {
				cmd.Println(stat)
			}
			return nil
		},
	}

	return cmd
}
