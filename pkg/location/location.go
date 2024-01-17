package location

type Heading byte

const (
	N Heading = 'N'
	W Heading = 'W'
	S Heading = 'S'
	E Heading = 'E'
)

type Movement byte

const (
	L Movement = 'L'
	R Movement = 'R'
	M Movement = 'M'
)

type Coordinates struct {
	X uint8
	Y uint8
}

type Position struct {
	Coordinates
	Heading
}

type Plateau struct {
	Width  uint8
	Height uint8
}
