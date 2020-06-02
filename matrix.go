package matrix

import (
	"errors"
)

type Matrix struct {
	values [][]float64
}

func (m1 Matrix) rows() int {
	return len(m1.values)
}

func (m1 Matrix) cols() int {
	return len(m1.values[0])
}

func createEmptyMatrix(rows int, cols int) Matrix {
	values := make([][]float64, rows)

	for i := range values {
		values[i] = make([]float64, cols)
	}

	return Matrix{values}
}

func (m1 *Matrix) AddScalar(scalar float64) {
	for i, row := range m1.values {
		for j, _ := range row {
			m1.values[i][j] += scalar
		}
	}
}

func (m1 *Matrix) SubtractScalar(scalar float64) {
	m1.AddScalar(-scalar)
}

func (m1 *Matrix) MultiplyScalar(scalar float64) {
	for i, row := range m1.values {
		for j, _ := range row {
			m1.values[i][j] *= scalar
		}
	}
}

func (m1 *Matrix) AddMatrix(m2 Matrix) error {
	if m1.rows() != m2.rows() || m1.cols() != m2.cols() {
		return errors.New("Dimensions doesn't match")
	}

	for i, row := range m1.values {
		for j, _ := range row {
			m1.values[i][j] += m2.values[i][j]
		}
	}

	return nil
}

func (m1 *Matrix) SubtractMatrix(m2 Matrix) error {
	m2.MultiplyScalar(-1)
	return m1.AddMatrix(m2)
}

func (m1 Matrix) MultiplyMatrix(m2 Matrix) (Matrix, error) {
	if m1.cols() != m2.rows() {
		return Matrix{}, errors.New("Dimension error")
	}

	result := createEmptyMatrix(m1.rows(), m2.cols())

	for i := 0; i < m1.rows(); i++ {
		for p := 0; p < result.cols(); p++ {
			for j := 0; j < m1.cols(); j++ {
				result.values[i][p] += m1.values[i][j] * m2.values[j][p]
			}
		}
	}

	return result, nil
}

func (m1 Matrix) Transpose() Matrix {
	result := createEmptyMatrix(m1.cols(), m1.rows())

	for i, row := range m1.values {
		for j, value := range row {
			result.values[j][i] = value
		}
	}

	return result
}

func (m1 Matrix) Trace() (float64, error) {
	if m1.rows() != m1.cols() {
		return 0, errors.New("Trace is only defined for nxn square matrices")
	}

	trace := 0.0

	for i := 0; i < m1.rows(); i++ {
		trace += m1.values[i][i]
	}

	return trace, nil
}
