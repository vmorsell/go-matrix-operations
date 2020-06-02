package matrix

import "math"

type Vector struct {
	components []float64
}

func (v1 Vector) Magnitude() float64 {
	componentsSquaredSum := 0.0
	for _, component := range v1.components {
		componentsSquaredSum += math.Pow(component, 2)
	}

	return math.Sqrt(componentsSquaredSum)
}

func (v1 *Vector) Normalize() {
	magnitude := v1.Magnitude()

	for i := range v1.components {
		v1.components[i] /= magnitude
	}
}

func (v1 *Vector) AddVector(v2 Vector) {
	for i := range v1.components {
		v1.components[i] += v2.components[i]
	}
}

func (v1 *Vector) SubtractVector(v2 Vector) {
	v2.MultiplyScalar(-1)
	v1.AddVector(v2)
}

func (v1 *Vector) MultiplyScalar(scalar float64) {
	for i := range v1.components {
		v1.components[i] *= scalar
	}
}
