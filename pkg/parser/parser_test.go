package parser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	. "github.com/czernous/mars-rover/pkg/location"
	. "github.com/czernous/mars-rover/pkg/rover"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      []MarsRover
		expectedError string
	}{
		{
			name: "Valid input",
			input: `
			    5 5

			    1 2 N

			    LMLMLMLMM

			    3 3 E
				
			    MMRMMRMRRM
			`,

			expected: []MarsRover{
				*NewMarsRover(
					Plateau{
						Width:  5,
						Height: 5,
					},
					Position{
						Coordinates: Coordinates{
							X: 1,
							Y: 2,
						},
						Heading: 'N',
					}, "LMLMLMLMM"),
				*NewMarsRover(
					Plateau{
						Width:  5,
						Height: 5,
					}, Position{
						Coordinates: Coordinates{
							X: 3,
							Y: 3,
						},
						Heading: 'E',
					}, "MMRMMRMRRM"),
			},
		},
		{
			name: "Valid input",
			input: `
				4 4
			    Asdasdgbwg
			    12N
				LMLMLMLMM
				9999
			    3 3 E
				
			    MMRMMRMRRM
			`,

			expected: []MarsRover{
				*NewMarsRover(
					Plateau{
						Width:  4,
						Height: 4,
					},
					Position{
						Coordinates: Coordinates{
							X: 3,
							Y: 3,
						},
						Heading: 'E',
					},

					"LMLMLMLMM"),
			},
		},
		{
			name: "Multiple rovers and commands",
			input: `
				5 5
				1 2 N
				LMLMLMLMM
				3 3 E
				MMRMMRMRRM
				2 2 W
				MRLMRMM
			`,
			expected: []MarsRover{
				*NewMarsRover(
					Plateau{
						Width:  5,
						Height: 5,
					},
					Position{
						Coordinates: Coordinates{
							X: 1,
							Y: 2,
						},
						Heading: 'N',
					},
					"LMLMLMLMM"),
				*NewMarsRover(
					Plateau{
						Width:  5,
						Height: 5,
					},
					Position{
						Coordinates: Coordinates{
							X: 3,
							Y: 3,
						},
						Heading: 'E',
					},
					"MMRMMRMRRM"),
				*NewMarsRover(
					Plateau{
						Width:  5,
						Height: 5,
					},
					Position{
						Coordinates: Coordinates{
							X: 2,
							Y: 2,
						},
						Heading: 'W',
					},
					"MRLMRMM"),
			},
		},
		{
			name: "Invalid plateau format",
			input: `
				A B
				1 2 N
				LMLMLMLMM
			`,
			expectedError: "Error: input contains no Plateau dimensions",
		},
		{
			name: "Invalid plateau coordinates",
			input: `
				0 5
				1 2 N
				LMLMLMLMM
			`,
			expectedError: "Error: input contains no Plateau dimensions", // plateau is not set if h or w is less than 3
		},
		{
			name: "Invalid position format",
			input: `
				5 5
				A B N
				LMLMLMLMM
			`,
			expectedError: "Error: input contains no valid rover movement data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes := strings.NewReader(tt.input)

			rovers, err := ParseInput(bytes)
			if tt.expectedError != "" {

				if err == nil || err.Error() != tt.expectedError {
					t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
				}
			} else {

				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				if len(rovers) != len(tt.expected) {
					t.Errorf("Wrong number of rovers, want = %d, got = %d", len(tt.expected), len(rovers))
				}

				for i := 0; i < len(rovers); i++ {

					if !reflect.DeepEqual(rovers[i], tt.expected[i]) {
						t.Errorf("Rover does not match test data, want = %v, got  %v", tt.expected[i], rovers[i])
					}
				}
			}

		})
	}
}

func TestParsePosition(t *testing.T) {
	input := "5 2 S"

	expected := &Position{
		Coordinates: Coordinates{
			X: 5,
			Y: 2,
		},
		Heading: 'S',
	}

	p, err := ParsePosition(input)

	if err != nil {
		fmt.Printf("Error parsing position: %v", err.Error())
	}

	if !reflect.DeepEqual(&p, expected) {
		t.Errorf("Incorrect position, got = %v, want = %v", &p, expected)
	}
}

func TestParseInvalidPosition(t *testing.T) {
	invalidInputs := []string{
		"invalid format",
		"1 2 X",
		"10 5",
		"20 S 30",
	}

	for _, input := range invalidInputs {
		_, err := ParsePosition(input)

		if err == nil {
			t.Errorf("No error was returned for incorrect input: %s", input)
		}
	}
}
