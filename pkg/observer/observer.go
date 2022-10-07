package observer

import (
	"math"
)

type Vector struct {
	X, Y float64
}

type Observer struct {
	Position  Vector
	Direction Vector
	Plane     Vector
}

func (o *Observer) Rotate(angle float64) {
	var dirMatrix [2]float64 = [2]float64{
		o.Direction.X, o.Direction.Y,
	}
	var planeMatrix [2]float64 = [2]float64{
		o.Plane.X, o.Plane.Y,
	}
	var c [2]float64
	var d [2]float64
	rotationMatrix := [2][2]float64{
		{math.Cos(angle), -math.Sin(angle)},
		{math.Sin(angle), math.Cos(angle)},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i] += dirMatrix[j] * rotationMatrix[i][j]
			d[i] += planeMatrix[j] * rotationMatrix[i][j]
		}
	}
	o.Direction.X, o.Direction.Y = c[0], c[1]
	o.Plane.X, o.Plane.Y = d[0], d[1]
}
