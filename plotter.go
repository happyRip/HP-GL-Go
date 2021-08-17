package plotter

type Pen struct {
	x, y int
}

func NewPen(x, y float64) Pen {
	p := Pen{}
	return p
}

func (p Pen) Up(coordinates ...float64) string {
	s := "PU"

	if len(coordinates)%2 != 0 {
		coordinates = append(coordinates, 0)
	}

	for i := 0; i < len(coordinates); i += 2 {
		x, y := coordinates[i], coordinates[i+1]
		s += string(x) + "," + string(y)
		if i < len(coordinates)-2 {
			s += ","
		}
	}

	s += ";"
	return s
}

func (p *Pen) SetX(f float64) {
	p.x = floatToUnit(f)
}

func (p *Pen) SetY(f float64) {
	p.y = floatToUnit(f)
}

func (p Pen) X() float64 {
	return unitToFloat(p.x)
}

func (p Pen) Y() float64 {
	return unitToFloat(p.y)
}
