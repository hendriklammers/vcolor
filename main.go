package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		// fmt.Fprintln(os.Stderr, "Please provide a valid hex color")
		// os.Exit(1)
		log.Fatal("Please provide a valid hex color")
	}

	fmt.Println(parseColor(args[1]))
}

func parseColor(hex string) string {
	re := regexp.MustCompile("(?i)^#?([a-f0-9]{2})([a-f0-9]{2})([a-f0-9]{2})$")
	values := make([]string, 3)

	for index, match := range re.FindStringSubmatch(hex)[1:] {
		num, err := strconv.ParseInt(match, 16, 16)
		if err != nil {
			log.Fatal(err)
		}

		val := fmt.Sprintf("%.3f", float32(num)/255)

		values[index] = stripZeros(val)
	}

	return fmt.Sprintf("vec3(%v)", strings.Join(values, ", "))
}

func stripZeros(str string) string {
	re := regexp.MustCompile("0{0,2}$")
	return re.ReplaceAllString(str, "")
}
