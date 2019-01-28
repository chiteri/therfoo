package tensor

func Mean(v *Vector) float64 {
	a := 0.
	v.Each(func(index int, value float64) {
		a += value
	})
	return a / float64(len(*v))
}
