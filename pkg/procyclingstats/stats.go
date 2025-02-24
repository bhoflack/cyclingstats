package procyclingstats

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bhoflack/cyclingstats/pkg/db"
	"github.com/bhoflack/cyclingstats/pkg/model"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Date     string
	Result   int
	Race     string
	Distance float32
}

type Stats struct {
	UpcomingRaces []string
	Results       []Result
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

	var results []Result

	doc.Find("table.rdrResults tbody tr").Each(func(i int, s *goquery.Selection) {
		var result Result
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				result.Date = s.Text()
			case 1:
				result.Result = -1
				if s.Text() != "DNF" && s.Text() != "DNS" && s.Text() != "OTL" && s.Text() != "NR" && s.Text() != "DSQ" && s.Text() != "" {
					v, err := strconv.Atoi(s.Text())
					if err != nil {
						log.Fatalf("error converting %s to int: %v", s.Text(), err)
					}
					result.Result = v
				}
			case 4:
				result.Race = strings.Trim(s.Find("a").Contents().First().Text(), " \n\r")
			case 5:

				if s.Text() != "" {
					distance, err := strconv.ParseFloat(s.Text(), 32)
					if err != nil {
						log.Printf("error converting %s to float: %v", s.Text(), err)
					}
					result.Distance = float32(distance)
				}
			}
		})
		results = append(results, result)
	})

	return Stats{UpcomingRaces: upcomingRaces, Results: results}, nil
}
