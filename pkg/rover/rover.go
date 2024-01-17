package rover

import (
	"errors"

	. "github.com/czernous/mars-rover/pkg/location"
)

type Rover interface {
	GetPosition() Position

	SetPosition(x uint8, y uint8, h Heading)

	GetHeading() Heading

	SetHeading(h Heading)

	Advance()

	Turn(direction byte)
}

type MarsRover struct {
	position Position
	Commands string
	Plateau
}

func (mr *MarsRover) GetPosition() Position {
	return mr.position
}

func (mr *MarsRover) SetPosition(x uint8, y uint8, h Heading) {
	mr.position.X = x
	mr.position.Y = y
	mr.position.Heading = h
}

func (mr *MarsRover) GetHeading() Heading {
	return mr.position.Heading
}

func (mr *MarsRover) SetHeading(h Heading) {
	mr.position.Heading = h
}

func (mr *MarsRover) SetPlateau(pl Plateau) {
	mr.Plateau = pl
}

func (mr *MarsRover) Advance(pl *Plateau) error {
	p := mr.position

	var x, y uint8

	switch mr.position.Heading {
	case N:
		x, y = p.X, p.Y+1
	case S:
		x, y = p.X, p.Y-1
	case E:
		x, y = p.X+1, p.Y
	case W:
		x, y = p.X-1, p.Y
	}

	if x >= 0 && x <= pl.Width && y >= 0 && y <= pl.Height {
		mr.SetPosition(x, y, mr.position.Heading)
		return nil
	} else {
		return errors.New("Rover cannot move beyond the plateau boundaries.")
	}
}

func (mr *MarsRover) Turn(d byte) {
	switch d {
	case 'L':
		mr.turnLeft()
	case 'R':
		mr.turnRight()
	}
}

func (mr *MarsRover) turnLeft() {
	switch mr.position.Heading {
	case N:
		mr.SetHeading(W)
	case W:
		mr.SetHeading(S)
	case S:
		mr.SetHeading(E)
	case E:
		mr.SetHeading(N)
	}
}

func (mr *MarsRover) turnRight() {
	switch mr.position.Heading {
	case N:
		mr.SetHeading(E)
	case E:
		mr.SetHeading(S)
	case S:
		mr.SetHeading(W)
	case W:
		mr.SetHeading(N)
	}
}

func NewMarsRover(pl Plateau, p Position, c string) *MarsRover {
	mr := &MarsRover{
		Commands: c,
		Plateau:  pl,
	}

	mr.SetPosition(p.X, p.Y, p.Heading)
	return mr
}
