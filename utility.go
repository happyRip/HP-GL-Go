package plotter

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const UNIT = 40 // HPGL uses resolution of 40 units per mm

func floatToUnit(f float64) int {
	return int(math.Round(f * UNIT))
}

func floatToUnitString(f float64) string {
	return fmt.Sprint(floatToUnit(f))
}

func floatToIntTimesTen(f float64) int {
	return int(math.Round(f * 10))
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func floatToSlice(f ...float64) []float64 {
	return f
}

func floatToUnitSlice(f ...float64) []int {
	var output []int
	for _, v := range f {
		output = append(output, floatToUnit(v))
	}
	return output
}

func unitToFloat(i int) float64 {
	return float64(i / UNIT)
}

func intSingleDecimalToFloat(i int) float64 {
	return float64(i) / 10
}

func toStringUnits(i int) string {
	var f float64 = intSingleDecimalToFloat(i)
	return strconv.FormatFloat(f*UNIT, 'f', -1, 64)
}

func getNumbers(s string) []string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	return re.FindAllString(s, -1)
}

func GetDimensionsFromFile(source string) (floatPair, error) {
	empty := floatPair{}

	if extension := filepath.Ext(source); extension != ".plt" {
		return empty, errors.New("incorrect file type")
	}

	file, err := os.Open(source)
	if err != nil {
		return empty, err
	}

	x, y := extremes{}, extremes{}
	x.init()
	y.init()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[:2] == "PD" {
			stringSlice := getNumbers(scanner.Text())

			for i, v := range stringSlice {
				v, err := strconv.Atoi(v)
				if err != nil {
					return empty, err
				}

				switch i % 2 {
				case 0:
					x.setExtremes(v)
				case 1:
					y.setExtremes(v)
				}

			}
		}
	}

	dimensions := floatPair{
		x: float64(x.max-x.min) / UNIT,
		y: float64(y.max-y.min) / UNIT,
	}

	err = file.Close()
	if err != nil {
		return empty, err
	}
	return dimensions, nil
}

type extremes struct {
	min, max int
}

func (e *extremes) init() {
	e.min, e.max = math.MaxInt64, math.MinInt64
}

func (e *extremes) setExtremes(i int) {
	e.setMin(i)
	e.setMax(i)
}

func (e *extremes) setMin(i int) {
	if e.min > i {
		e.min = i
	}
}

func (e *extremes) setMax(i int) {
	if e.max < i {
		e.max = i
	}
}

type floatPair struct {
	x, y float64
}
