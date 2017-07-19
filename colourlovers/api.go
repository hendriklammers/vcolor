package colourlovers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Color is used to store color json
type Color struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Hex         string `json:"hex"`
}

// Palette is used to store color palette json
type Palette struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Colors      []string `json:"colors"`
}

// GetRandomColor gets a random color from the Colourlovers API
func GetRandomColor() (string, error) {
	url := "http://www.colourlovers.com/api/colors/random?format=json"
	return getColorData(url)
}

// GetPalette gets colors of palette with provided id from the Colourlovers API
func GetPalette(id int) ([]string, error) {
	url := fmt.Sprintf("http://www.colourlovers.com/api/palette/%d?format=json", id)
	return getPaletteData(url)
}

// GetRandomPalette gets a random color palette from the Colourlovers API
func GetRandomPalette() ([]string, error) {
	url := "http://www.colourlovers.com/api/palettes/random?format=json"
	return getPaletteData(url)
}

func getColorData(url string) (string, error) {
	var data []Color
	err := getJSON(url, &data)
	if err != nil {
		return "", err
	}
	if len(data) < 1 {
		return "", errors.New("No color returned from Colourlovers API")
	}
	return data[0].Hex, nil
}

func getPaletteData(url string) ([]string, error) {
	var data []Palette
	err := getJSON(url, &data)
	if err != nil {
		return nil, err
	}
	if len(data) < 1 {
		return nil, errors.New("No color palette returned from Colourlovers API")
	}
	return data[0].Colors, nil
}

func getJSON(url string, target interface{}) error {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
