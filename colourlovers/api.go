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

	data, err := getJSON(url)
	if err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, errors.New("No palette found for provided ID")
	}
	return data[0].Colors, nil
}

// RandomPalette gets a random color palette Colourlovers API
func RandomPalette() ([]string, error) {
	url := "http://www.colourlovers.com/api/palettes/random?format=json"

	data, err := getJSON(url)
	if err != nil {
		return nil, err
	}

	if len(data) < 1 {
		return nil, errors.New("No random color palette returned from Colourlovers API")
	}
	return data[0].Colors, nil
}

func getJSON(url string) (response, error) {
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
	return data, nil
}
