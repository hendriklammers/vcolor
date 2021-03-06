package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hendriklammers/vcolor/colourlovers"
)

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

	for _, color := range getColors() {
		hex, ok := colorNames[strings.ToLower(color)]
		if !ok {
			hex = color
		}
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
