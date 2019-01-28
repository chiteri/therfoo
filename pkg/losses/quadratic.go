package losses

import (
	"github.com/therfoo/therfoo/pkg/tensor"
	"math"
)

func Quadratic(yTrue, yEstimate *tensor.Vector) *tensor.Vector {
	e := tensor.Vector{}
	yEstimate.Each(func(index int, value float64) {
		e.Append(math.Pow(value-(*yTrue)[index], 2))
	})
	return &e
}

func QuadraticPrime(yTrue, yEstimate *tensor.Vector) *tensor.Vector {
	e := tensor.Vector{}
	yEstimate.Each(func(index int, value float64) {
		e.Append(2. * (value - (*yTrue)[index]))
	})
	return &e
}
