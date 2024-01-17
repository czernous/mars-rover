package parser

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

	. "github.com/czernous/mars-rover/pkg/location"
	. "github.com/czernous/mars-rover/pkg/rover"
)

var validPlateauRegex = regexp.MustCompile(`^\d\s\d$`)
var validPositionRegex = regexp.MustCompile(`^\d\s\d\s[NSWE]$`)
var validCommandsRegex = regexp.MustCompile(`^[LRM]*$`)

/**
* ParseInput scans input line by line, looks for Position and Commands.
* Adds first found match to the respective Rover field, keeps scanning until the other field is matched
* If either Position or Command field is already assigned, next found match is skipped... see test file for
 */

func ParseInput(i io.Reader) ([]MarsRover, error) {
	scanner := bufio.NewScanner(i)

	roverQueue := []MarsRover{}
	var r *MarsRover
	var plateau Plateau

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if validPlateauRegex.MatchString(line) {
			if plateau == (Plateau{}) {
				plateau, _ = ParsePlateau(line)
			}
		}

		if validPositionRegex.MatchString(line) {
			if r == nil {
				r = &MarsRover{Commands: ""}
			}
			p, _ := ParsePosition(line)
			r.SetPosition(p.X, p.Y, p.Heading)
		}

		if validCommandsRegex.MatchString(line) {
			if r == nil {
				r = &MarsRover{}
			}
			r.Commands = line
		}

		if r != nil && r.Commands != "" && r.GetPosition() != (Position{}) && plateau != (Plateau{}) {
			r.SetPlateau(plateau)
			roverQueue = append(roverQueue, *r)
			r = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return roverQueue, errors.New("Error reading input: " + err.Error())
	}

	if plateau == (Plateau{}) {
		return nil, errors.New("Error: input contains no Plateau dimensions")
	}

	if len(roverQueue) < 1 {
		return roverQueue, errors.New("Error: input contains no valid rover movement data")
	}

	return roverQueue, nil
}

func ParsePosition(i string) (Position, error) {
	posArr := strings.Fields(i)

	if len(posArr) != 3 {
		return Position{}, errors.New("Invalid position format")
	}

	x, err := strconv.ParseUint(posArr[0], 10, 8)
	if err != nil {
		return Position{}, errors.New("Invalid X coordinate")
	}

	y, err := strconv.ParseUint(posArr[1], 10, 8)
	if err != nil {
		return Position{}, errors.New("Invalid Y coordinate")
	}

	h := Heading(posArr[2][0])

	if h != 'N' && h != 'S' && h != 'W' && h != 'E' {
		return Position{}, errors.New("Invalid Heading format")
	}

	return Position{
		Coordinates: Coordinates{
			X: uint8(x),
			Y: uint8(y),
		},
		Heading: h,
	}, nil
}

func ParsePlateau(i string) (Plateau, error) {
	pArr := strings.Fields(i)

	if len(pArr) != 2 {
		return Plateau{}, errors.New("Invalid plateau format")
	}

	width, err := strconv.ParseUint(pArr[0], 10, 8)
	if err != nil || width < 3 {
		return Plateau{}, errors.New("Invalid Width")
	}

	height, err := strconv.ParseUint(pArr[1], 10, 8)
	if err != nil || height < 3 {
		return Plateau{}, errors.New("Invalid Height")
	}

	return Plateau{
		Width:  uint8(width),
		Height: uint8(height),
	}, nil

}
