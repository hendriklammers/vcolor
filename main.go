package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hendriklammers/vcolor/colourlovers"
)

const regex = "(?i)^#?([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})?$"

var (
	paletteFlag = flag.Int(
		"p",
		0,
		"get a palette by ID from the Colourlovers API",
	)

	randomPaletteFlag = flag.Bool(
		"rp",
		false,
		"get a random palette from the Colourlovers API",
	)

	randomColorFlag = flag.Bool(
		"r",
		false,
		"get a random color from the Colourlovers API",
	)
)

func main() {
	flag.Parse()

	for _, hex := range getColors() {
		vec, err := ParseColor(hex)
		if err != nil {
			// Maybe should just print error instead of exit?
			log.Fatal(err)
		}
		fmt.Println(vec)
	}
}

func getColors() []string {
	var (
		colors []string
		err    error
	)

	if *randomColorFlag {
		var color string
		color, err = colourlovers.GetRandomColor()
		if err != nil {
			log.Fatal(err)
		}
		colors = append(colors, color)
	} else if *paletteFlag > 0 {
		colors, err = colourlovers.GetPalette(*paletteFlag)
		if err != nil {
			log.Fatal(err)
		}
	} else if *randomPaletteFlag {
		colors, err = colourlovers.GetRandomPalette()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		args := os.Args
		if len(args) < 2 {
			fmt.Printf("usage: %s [hex-color]\n", os.Args[0])
			os.Exit(1)
		}
		colors = args[1:]
	}

	return colors
}

// ParseColor takes hex color string and converts it to a string in vec3 or vec4 format
func ParseColor(hex string) (string, error) {
	re := regexp.MustCompile(regex)
	if !re.MatchString(hex) {
		msg := fmt.Sprintf("%s is not a valid hex color", hex)
		return "", errors.New(msg)
	}

	var values []string
	for _, match := range re.FindStringSubmatch(hex)[1:] {
		if match != "" {
			num, err := strconv.ParseInt(match, 16, 16)
			if err != nil {
				return "", err
			}
			val := fmt.Sprintf("%.3f", float32(num)/255)
			values = append(values, stripZeros(val))
		}
	}

	return fmt.Sprintf("vec%v(%v)", len(values), strings.Join(values, ", ")), nil
}

func stripZeros(str string) string {
	re := regexp.MustCompile("0{0,2}$")
	return re.ReplaceAllString(str, "")
}
