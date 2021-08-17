package plotter

import "testing"

func TestUp(t *testing.T) {
	SetCommandSeparator("\n")
	coordinates := [][]float64{
		{},
		{0, 1},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2},
	}
	want := []string{
		"PU;\n",
		"PU0,40;\n",
		"PU0,40,80,120,160,200;\n",
		"PU0,40,80,0;\n",
	}
	p := NewPen(6, 7)
	for i, c := range coordinates {
		got := p.Up(c...)
		position := floatToSlice(p.Position())
		if want[i] != got {
			t.Error(
				"\nwant:", want[i],
				"\ngot:", got,
				"\nposition:", position,
			)
		}
	}
}
