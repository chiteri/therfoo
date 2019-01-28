package tensor

import (
	"fmt"
	"testing"
)

func TestScalarEach(t *testing.T) {
	tests := []float64{1., 2., 3., 5., 8., 13., 21., 34.}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%f", test), func(t *testing.T) {
			s := Scalar(test)
			actual := 0.
			s.Each(func(index int, value float64) {
				actual = value
			})
			if test != actual {
				t.Errorf("expected scalar to contain %f got %f", test, actual)
			}
		})
	}
}
