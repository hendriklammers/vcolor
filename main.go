package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const regex = "(?i)^#?([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})?$"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: %s [hex-color]\n", os.Args[0])
		os.Exit(1)
	}

	for _, hex := range args[1:] {
		vec, err := ParseColor(hex)
		if err != nil {
			// Maybe should just print error instead of exit?
			log.Fatal(err)
		}
		fmt.Println(vec)
	}
}

// ParseColor takes hex color string and converts it to a string in vec3 format
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
