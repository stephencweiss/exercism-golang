// Package triangle implements a simple utility to determine the type of a triangle given three sides
package triangle

import (
	"fmt"
	"math"
	"sort"
)

// Kind is the string return value from our const.
type Kind = string

const (
	// NaT is not a triangle
	NaT = "NaT"
	// Equ = equilateral
	Equ = "Equ"
	// Iso = isosceles
	Iso = "Iso"
	// Sca = scalene
	Sca = "Sca"
	// Deg = degenerate
	Deg = "Deg"
)

// KindFromSides takes three sides to determine the type of triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	fmt.Println((sides))
	triangleEquality := sides[0]+sides[1] >= sides[2] && sides[0] >= 0

	switch {
	case !triangleEquality || sides[0] == 0 || sides[2] == math.Inf(1):
		k = NaT
	case sides[0]+sides[1] == sides[2]:
		k = Deg
	case sides[0] == sides[1] && sides[0] == sides[2]:
		k = Equ
	case sides[1] == sides[2] && sides[0] < sides[1] && sides[0] < sides[2]:
		k = Iso
	case sides[0] != sides[1] && sides[0] != sides[2]:
		k = Sca
	default:
		panic(fmt.Sprintf("Unknown case! %f, %f, %f", a, b, c))
	}

	return k
}
