package main

import (
	"log"
	"testing"
)

type ParseColorTestPair struct {
	input  string
	output string
}

func TestParseColor(t *testing.T) {
	tests := []ParseColorTestPair{
		{"ffffff", "vec3(1.0, 1.0, 1.0)"},
		{"000000", "vec3(0.0, 0.0, 0.0)"},
		{"ABC123", "vec3(0.671, 0.757, 0.137)"},
		{"#ff6600ff", "vec4(1.0, 0.4, 0.0, 1.0)"},
	}

	for _, test := range tests {
		result, err := ParseColor(test.input)
		if err != nil {
			log.Fatal(err)
		}
		if result != test.output {
			t.Error(
				"For input:", test.input,
				"expected:", test.output,
				"got:", result,
			)
		}
	}
}
