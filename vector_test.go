package matrix

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func (v1 Vector) equals(v2 Vector) error {
	if len(v1.components) != len(v2.components) {
		return errors.New(fmt.Sprintf("Non matching number of components. v1 has %d, v2 has %d", len(v1.components), len(v2.components)))
	}

	for i := range v1.components {
		if v1.components[i] != v2.components[i] {
			return errors.New(fmt.Sprintf("Non-matching component at position %d. v1 has %f, v2 has %f", i, v1.components[i], v2.components[i]))
		}
	}

	return nil
}

func vectorInOtherVectorSpace(v1 Vector) Vector {
	tooManyComponents := make([]float64, len(v1.components)+1)

	return Vector{tooManyComponents}
}

func TestVectorMagnitude(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	correct := math.Sqrt(83)

	magnitude := v1.Magnitude()
	if magnitude != correct {
		t.Fatalf("Unexpected result %f, expected %f.", magnitude, correct)
	}
}

func TestVectorNormalize(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	correct := Vector{[]float64{-3 / math.Sqrt(83), 0, 5 / math.Sqrt(83), 7 / math.Sqrt(83)}}

	v1.Normalize()

	error := v1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestVectorAddVector(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	v2 := Vector{[]float64{5, 3, -1, -10}}
	correct := Vector{[]float64{2, 3, 4, -3}}

	v1.AddVector(v2)

	error := v1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestVectorSubtractVector(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	v2 := Vector{[]float64{5, 3, -1, -10}}
	correct := Vector{[]float64{-8, -3, 6, 17}}

	v1.SubtractVector(v2)

	error := v1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestVectorMultiplyScalar(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	scalar := -3.0
	correct := Vector{[]float64{9, 0, -15, -21}}

	v1.MultiplyScalar(scalar)

	error := v1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestVectorDotProduct(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	v2 := Vector{[]float64{9, 0, -15, -21}}

	dotProduct, error := v1.DotProduct(vectorInOtherVectorSpace(v1))
	if error == nil {
		t.Fatalf("Expected error. Number of components doesn't match")
	}

	dotProduct, error = v1.DotProduct(v2)
	if dotProduct != -249 {
		t.Fatalf("Unexpected return. Expected %d, got %f", -249, dotProduct)
	}
	if error != nil {
		t.Fatalf("Unexpected error.")
	}
}

func TestVectorDistance(t *testing.T) {
	v1 := Vector{[]float64{-3, 0, 5, 7}}
	v2 := Vector{[]float64{9, 0, -15, -21}}
	correct := 4 * math.Sqrt(83)

	_, error := v1.Distance(vectorInOtherVectorSpace(v1))
	if error == nil {
		t.Fatalf("Expected error. Number of components doesn't match")
	}

	distance, error := v1.Distance(v2)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
	if distance != correct {
		t.Fatalf("Unexpected return. Got %f, expected %f", distance, correct)
	}
}
