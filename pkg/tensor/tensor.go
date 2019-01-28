package tensor

type Tensor interface {
	Each(func(index int, value float64))
}
