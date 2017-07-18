package colourlovers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type palette struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Colors      []string `json:"colors"`
}

type response []palette

// Palette gets colors of palette with provided id from the Colourlovers API
func Palette(id int) ([]string, error) {
	url := fmt.Sprintf("http://www.colourlovers.com/api/palette/%d?format=json", id)

	// Using custom http client so a timeout can be set
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data := response{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, errors.New("No palette found for provided ID")
	}
	return data[0].Colors, nil
}

// RandomPalette gets colors of a random palette from the Colourlovers API
func RandomPalette() ([]string, error) {
	return nil, nil
}
