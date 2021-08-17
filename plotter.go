package plotter

import (
	"fmt"
)

var commandSeparator string

func SetCommandSeparator(separator string) {
	commandSeparator = separator
}

func CommandSeparator() string {
	return commandSeparator
}

type Mode int

const (
	Absolute Mode = iota // Absolute is a default value here
	Relative
)

type Pen struct {
	x, y int
}

func NewPen(x, y float64) Pen {
	p := Pen{}
	return p
}

func (p *Pen) Up(coordinates ...float64) string {
	if len(coordinates) >= 2 {
		if len(coordinates)%2 != 0 {
			coordinates = append(coordinates, 0)
		}
		p.SetPosition(
			coordinates[len(coordinates)-2],
			coordinates[len(coordinates)-1],
		)
	}
	return moveCommand("PU", coordinates...)
}

func (p *Pen) Down(coordinates ...float64) string {
	if len(coordinates) >= 2 {
		if len(coordinates)%2 != 0 {
			coordinates = append(coordinates, 0)
		}
		p.SetPosition(
			coordinates[len(coordinates)-2],
			coordinates[len(coordinates)-1],
		)
	}
	return moveCommand("PD", coordinates...)
}

func (p *Pen) Absolute(coordinates ...float64) string {
	if len(coordinates) >= 2 {
		if len(coordinates)%2 != 0 {
			coordinates = append(coordinates, 0)
		}
		p.SetPosition(
			coordinates[len(coordinates)-2],
			coordinates[len(coordinates)-1],
		)
	}
	return moveCommand("PA", coordinates...)
}

func (p *Pen) Relative(coordinates ...float64) string {
	if len(coordinates) >= 2 {
		if len(coordinates)%2 != 0 {
			coordinates = append(coordinates, 0)
		}
		p.SetPosition(
			p.X()+coordinates[len(coordinates)-2],
			p.Y()+coordinates[len(coordinates)-1],
		)
	}
	return moveCommand("PR", coordinates...)
}

func (p *Pen) Move(mode Mode, coordinates ...float64) string {
	output := "PU;" + CommandSeparator()
	switch mode {
	case Absolute:
		output += p.Absolute(coordinates...)
	case Relative:
		output += p.Relative(coordinates...)
	}
	return output
}

func (p *Pen) Line(mode Mode, coordinates ...float64) string {
	output := "PD;" + CommandSeparator()
	switch mode {
	case Absolute:
		output += p.Absolute(coordinates...)
	case Relative:
		output += p.Relative(coordinates...)
	}
	return output
}

func (p *Pen) SetX(f float64) {
	p.x = floatToUnit(f)
}

func (p *Pen) SetY(f float64) {
	p.y = floatToUnit(f)
}

func (p *Pen) SetPosition(x, y float64) {
	p.SetX(x)
	p.SetY(y)
}

func (p Pen) X() float64 {
	return unitToFloat(p.x)
}

func (p Pen) Y() float64 {
	return unitToFloat(p.y)
}

func (p Pen) Position() (float64, float64) {
	return p.X(), p.Y()
}

func ConstructCommand(command string, args ...int) string {
	for i, v := range args {
		command += fmt.Sprint(v)
		if i < len(args)-1 {
			command += ","
		}
	}
	command += ";" + commandSeparator
	return command
}

func moveCommand(command string, coordinates ...float64) string {
	unitCoordinates := floatToUnitSlice(coordinates...)
	s := ConstructCommand("PU", unitCoordinates...)
	return s
}
