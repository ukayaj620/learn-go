package cylinder

import (
	"math"
)

type Cylinder struct {
	Height float64
	Radius float64
}

func (cylinder Cylinder) SurfaceArea() float64 {
	return 2.0 * math.Pi * cylinder.Radius * (cylinder.Radius + cylinder.Height)
}

func (cylinder Cylinder) Volume() float64 {
	return math.Pi * math.Pow(cylinder.Radius, 2) * cylinder.Height
}
