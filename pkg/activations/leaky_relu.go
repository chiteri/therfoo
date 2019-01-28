package activations

import (
	"github.com/therfoo/therfoo/pkg/tensor"
)

func LeakyReLU(z *tensor.Vector) *tensor.Vector {
	a := tensor.Vector{}

	a.Each(func(index int, value float64) {
		if value >= 0 {
			a.Append(value)
		} else {
			a.Append(value / 100.)
		}
	})

	return &a
}

func LeakyReLUPrime(z *tensor.Vector) *tensor.Vector {
	a := tensor.Vector{}

	z.Each(func(index int, value float64) {
		if value > 0 {
			a.Append(1)
		} else {
			// TODO:
			a.Append(0)
		}
	})

	return &a
}
