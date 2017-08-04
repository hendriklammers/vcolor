package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const regex = "(?i)^#?([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})?$"

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
