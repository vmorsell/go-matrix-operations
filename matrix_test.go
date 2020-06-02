package matrix

import (
	"errors"
	"fmt"
	"testing"
)

func (m1 Matrix) equals(m2 Matrix) error {
	if len(m1.values) != len(m2.values) {
		return errors.New(fmt.Sprintf("Non-matching row dimensions. m1 has %d, m2 has %d.", len(m1.values), len(m2.values)))
	}

	if len(m1.values[0]) != len(m2.values[0]) {
		return errors.New(fmt.Sprintf("Non-matching column dimensions. m1 has %d, m2 has %d.", len(m1.values[0]), len(m2.values[0])))
	}

	for i, rows := range m1.values {
		for j, _ := range rows {
			if m1.values[i][j] != m2.values[i][j] {
				return errors.New(fmt.Sprintf("Values at position [%d][%d] doesn't match. m1 has %f, m2 has %f.", i, j, m1.values[i][j], m2.values[i][j]))
			}
		}
	}

	return nil
}

func TestMatrixDimensions(t *testing.T) {
	m1 := Matrix{[][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}}

	if m1.dim().rows != 2 {
		t.Fatalf("Unexpected return. Rows dimension method returned %d, expected %d.", m1.dim().rows, 2)
	}

	if m1.dim().cols != 4 {
		t.Fatalf("Unexpected return. Cols dimension method returned %d, expected %d.", m1.dim().cols, 4)
	}
}

func TestMatrixAddScalar(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	scalar := 1.5
	correct := Matrix{[][]float64{
		{-8.5, -3.5, 1.5, 6.5},
		{-3.5, 1.5, 6.5, 11.5},
	}}

	m1.AddScalar(scalar)

	error := m1.equals(correct)

	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixSubtractScalar(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	scalar := 1.5
	correct := Matrix{[][]float64{
		{-11.5, -6.5, -1.5, 3.5},
		{-6.5, -1.5, 3.5, 8.5},
	}}

	m1.SubtractScalar(scalar)

	error := m1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixMultiplyScalar(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	scalar := 1.5
	correct := Matrix{[][]float64{
		{-15, -7.5, 0, 7.5},
		{-7.5, 0, 7.5, 15},
	}}

	m1.MultiplyScalar(scalar)

	error := m1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixAddMatrix(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	m2 := Matrix{[][]float64{
		{-5, 0, 5, 10},
		{-10, -5, 0, 5},
	}}
	correct := Matrix{[][]float64{
		{-15, -5, 5, 15},
		{-15, -5, 5, 15},
	}}

	error := m1.AddMatrix(m2)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}

	error = m1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixSubtractMatrix(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	m2 := Matrix{[][]float64{
		{-5, 0, 5, 10},
		{-10, -5, 0, 5},
	}}
	correct := Matrix{[][]float64{
		{-5, -5, -5, -5},
		{5, 5, 5, 5},
	}}

	error := m1.SubtractMatrix(m2)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}

	error = m1.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixMultiplyMatrix(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5},
		{-5, 0},
		{0, 5},
	}}
	m2 := Matrix{[][]float64{
		{-5, 0, 5, 10, 15},
		{5, 10, 15, 20, 25},
	}}
	correct := Matrix{[][]float64{
		{25, -50, -125, -200, -275},
		{25, 0, -25, -50, -75},
		{25, 50, 75, 100, 125},
	}}

	m3, error := m1.MultiplyMatrix(m2)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}

	error = m3.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixTranspose(t *testing.T) {
	m1 := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}
	correct := Matrix{[][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	}}

	m2 := m1.Transpose()

	error := m2.equals(correct)
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}
}

func TestMatrixTrace(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -9, -8},
		{-5, -4, -3},
		{1, 2, 3},
	}}
	correct := -11.0

	trace, error := m1.Trace()
	if error != nil {
		t.Fatalf("Unexpected error. %s", error)
	}

	if trace != correct {
		t.Fatalf("Unexpected return. Got %f, expected %f.", trace, correct)
	}
}
