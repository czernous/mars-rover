package rover

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/czernous/mars-rover/pkg/location"
)

func TestGetRoverPosition(t *testing.T) {
	mr := &MarsRover{
		position: Position{
			Coordinates: Coordinates{X: 4, Y: 5},
			Heading:     'N',
		},
		Commands: "LRRRLM",
	}

	p := mr.GetPosition()

	if !reflect.DeepEqual(p, mr.position) {
		t.Errorf("Unexpected position, got = %v, want = %v", p, mr.position)
	}
}

func TestRoverSetPosition(t *testing.T) {
	mr := &MarsRover{
		position: Position{
			Coordinates: Coordinates{X: 0, Y: 0},
			Heading:     'W',
		},
		Commands: "M",
	}

	mr.SetPosition(10, 2, 'S')

	expected := Position{
		Coordinates: Coordinates{
			X: 10,
			Y: 2,
		},
		Heading: 'S',
	}

	if !reflect.DeepEqual(expected, mr.position) {
		t.Errorf("Position was not correctly updated, got = %v, want = %v", expected, mr.position)
	}
}

func TestRoverAdvance(t *testing.T) {
	pl := &Plateau{
		Width:  15,
		Height: 15,
	}

	mr := &MarsRover{
		position: Position{
			Coordinates: Coordinates{X: 0, Y: 0},
			Heading:     'N',
		},
		Commands: "M",
	}

	expected := &MarsRover{
		position: Position{
			Coordinates: Coordinates{X: 0, Y: 1},
			Heading:     'N',
		},
		Commands: "M",
	}

	err := mr.Advance(pl)

	if err != nil {
		fmt.Printf("Error advancing rover: %v", err.Error())
	}

	if !reflect.DeepEqual(mr, expected) {
		t.Errorf("Incorrect rover data, got = %v, want = %v", mr, expected)
	}
}

func TestRoverAdvanceBeyondPlateau(t *testing.T) {
	pl := &Plateau{
		Width:  15,
		Height: 15,
	}

	mr := &MarsRover{
		position: Position{
			Coordinates: Coordinates{X: 0, Y: 15},
			Heading:     'N',
		},
		Commands: "M",
	}

	err := mr.Advance(pl)

	if err == nil {
		t.Errorf("No error was returned advancing out of bounds.")
	}
}

func TestRoverAdvanceDirection(t *testing.T) {
	pl := &Plateau{
		Width:  15,
		Height: 15,
	}

	tests := []struct {
		name     string
		initial  Position
		heading  Heading
		commands string
		expected Position
	}{
		{
			name: "Advance North",
			initial: Position{
				Coordinates: Coordinates{X: 5, Y: 5},
				Heading:     'N',
			},
			heading:  'N',
			commands: "M",
			expected: Position{Coordinates: Coordinates{X: 5, Y: 6}, Heading: 'N'},
		},
		{
			name: "Advance East",
			initial: Position{
				Coordinates: Coordinates{X: 5, Y: 5},
				Heading:     'E',
			},
			heading:  'E',
			commands: "M",
			expected: Position{Coordinates: Coordinates{X: 6, Y: 5}, Heading: 'E'},
		},
		{
			name: "Advance West",
			initial: Position{
				Coordinates: Coordinates{X: 5, Y: 5},
				Heading:     'W',
			},
			heading:  'W',
			commands: "M",
			expected: Position{Coordinates: Coordinates{X: 4, Y: 5}, Heading: 'W'},
		},
		{
			name: "Advance South",
			initial: Position{
				Coordinates: Coordinates{X: 5, Y: 5},
				Heading:     'S',
			},
			heading:  'S',
			commands: "M",
			expected: Position{Coordinates: Coordinates{X: 5, Y: 4}, Heading: 'S'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MarsRover{
				position: Position{
					Coordinates: tt.initial.Coordinates,
					Heading:     tt.initial.Heading,
				},
				Commands: tt.commands,
			}

			err := mr.Advance(pl)
			if err != nil {
				t.Fatalf("Error advancing rover: %v", err)
			}

			if !reflect.DeepEqual(mr.position, tt.expected) {
				t.Errorf("Incorrect rover position, got = %v, want = %v", mr.position, tt.expected)
			}
		})
	}
}

func TestRoverTurning(t *testing.T) {
	tests := []struct {
		name     string
		startPos Position
		commands string
		expected Position
	}{
		{
			name: "Turn Left from North",
			startPos: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     N,
			},
			commands: "L",
			expected: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     W,
			},
		},
		{
			name: "Turn Right from North",
			startPos: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     N,
			},
			commands: "R",
			expected: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     E,
			},
		},
		{
			name: "Turn Left from South",
			startPos: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     S,
			},
			commands: "L",
			expected: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     E,
			},
		},
		{
			name: "Turn Right from South",
			startPos: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     S,
			},
			commands: "R",
			expected: Position{
				Coordinates: Coordinates{X: 0, Y: 0},
				Heading:     W,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MarsRover{
				position: Position{
					Coordinates: tt.startPos.Coordinates,
					Heading:     tt.startPos.Heading,
				},
				Commands: tt.commands,
			}

			mr.Turn(tt.commands[0])

			if !reflect.DeepEqual(mr.position, tt.expected) {
				t.Errorf("Incorrect rover data after turning, got = %v, want = %v", mr.position, tt.expected)
			}
		})
	}
}
