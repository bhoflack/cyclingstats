package procyclingstats

import (
	"encoding/json"
	"net/http"

	"github.com/bhoflack/cyclingstats/pkg/model"
)

func Find(name string) ([]model.Cyclist, error) {
	resp, err := http.Get("https://www.procyclingstats.com/resources/search.php?term=" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cyclists []model.Cyclist
	if err := json.NewDecoder(resp.Body).Decode(&cyclists); err != nil {
		return nil, err
	}
	return cyclists, nil
}
