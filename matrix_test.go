package matrix

import (
	"errors"
	"fmt"
	"testing"
)

func equals(m1 Matrix, m2 Matrix) error {
	if len(m1.values) != len(m2.values) {
		return errors.New("Row dimension doesn't match")
	}

	if len(m1.values[0]) != len(m2.values[0]) {
		return errors.New("Column dimension doesn't match")
	}

	for i, rows := range m1.values {
		for j, _ := range rows {
			if m1.values[i][j] != m2.values[i][j] {
				return errors.New(fmt.Sprintf("Values at position [%d][%d] doesn't match", i, j))
			}
		}
	}

	return nil
}

func TestDimensions(t *testing.T) {
	m1 := Matrix{[][]float64{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}}

	if m1.rows() != 2 {
		t.Fatalf("Rows dimension method returned %d, expected %d", m1.rows(), 2)
	}

	if m1.cols() != 4 {
		t.Fatalf("Cols dimension method returned %d, expected %d", m1.cols(), 4)
	}
}

func TestAddScalar(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	scalar := 1.5
	correct := Matrix{[][]float64{
		{-8.5, -3.5, 1.5, 6.5},
		{-3.5, 1.5, 6.5, 11.5},
	}}

	m1 = AddScalar(m1, scalar)

	error := equals(m1, correct)

	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestSubtractScalar(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	scalar := 1.5
	correct := Matrix{[][]float64{
		{-11.5, -6.5, -1.5, 3.5},
		{-6.5, -1.5, 3.5, 8.5},
	}}

	m1 = SubtractScalar(m1, scalar)

	error := equals(m1, correct)
	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestMultiplyScalar(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -5, 0, 5},
		{-5, 0, 5, 10},
	}}
	scalar := 1.5
	correct := Matrix{[][]float64{
		{-15, -7.5, 0, 7.5},
		{-7.5, 0, 7.5, 15},
	}}

	m1 = MultiplyScalar(m1, scalar)

	error := equals(m1, correct)
	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestAddMatrix(t *testing.T) {
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

	m1, error := AddMatrix(m1, m2)
	if error != nil {
		t.Fatalf("%s", error)
	}

	error = equals(m1, correct)
	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestSubtractMatrix(t *testing.T) {
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

	m1, error := SubtractMatrix(m1, m2)
	if error != nil {
		t.Fatalf("%s", error)
	}

	error = equals(m1, correct)
	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestMultiplyMatrix(t *testing.T) {
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

	m1, error := MultiplyMatrix(m1, m2)
	if error != nil {
		t.Fatalf("%s", error)
	}

	error = equals(m1, correct)
	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestTranspose(t *testing.T) {
	m1 := Matrix{[][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}}
	correct := Matrix{[][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	}}

	m1 = Transpose(m1)

	error := equals(m1, correct)
	if error != nil {
		t.Fatalf("%s", error)
	}
}

func TestTrace(t *testing.T) {
	m1 := Matrix{[][]float64{
		{-10, -9, -8},
		{-5, -4, -3},
		{1, 2, 3},
	}}
	correct := -11.0

	trace, error := Trace(m1)
	if error != nil {
		t.Fatalf("%s", error)
	}

	if trace != correct {
		t.Fatalf("Trace result incorrect")
	}
}
