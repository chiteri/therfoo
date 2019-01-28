package tensor

type Scalar float64

func (s *Scalar) Each(f func(index int, value float64)) {
	f(0, float64(*s))
}
