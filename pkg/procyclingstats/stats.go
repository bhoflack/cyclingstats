package procyclingstats

import (
	"io"
	"net/http"
	"strings"

	"github.com/bhoflack/cyclingstats/pkg/db"
	"github.com/bhoflack/cyclingstats/pkg/model"

	"github.com/PuerkitoBio/goquery"
)

type Stats struct {
	UpcomingRaces []string
}

func UpdateStats(cyclists []model.Cyclist) error {
	for _, cyclist := range cyclists {
		resp, err := http.Get("https://www.procyclingstats.com/rider/" + cyclist.Id)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// copy the body to a string
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		client, err := db.NewClient()
		if err != nil {
			return err
		}

		// write the page to the db
		if err := client.AddPage(cyclist, string(body)); err != nil {
			return err
		}
	}
	return nil
}

func Extract(body string) (Stats, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return Stats{}, err
	}

	var upcomingRaces []string
	doc.Find("h3:contains('Upcoming participations')").Each(func(i int, s *goquery.Selection) {
		s.Parent().Find("span a").Each(func(i int, s *goquery.Selection) {
			if strings.HasPrefix(s.AttrOr("href", ""), "race/") {
				upcomingRaces = append(upcomingRaces, s.Text())
			}
		})
	})

	return Stats{UpcomingRaces: upcomingRaces}, nil
}
