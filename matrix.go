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

func AddScalar(m1 Matrix, scalar float64) Matrix {
	for i, row := range m1.values {
		for j, _ := range row {
			m1.values[i][j] += scalar
		}
	}

	return m1
}

func SubtractScalar(m1 Matrix, scalar float64) Matrix {
	return AddScalar(m1, -scalar)
}

func MultiplyScalar(m1 Matrix, scalar float64) Matrix {
	for i, row := range m1.values {
		for j, _ := range row {
			m1.values[i][j] *= scalar
		}
	}

	return m1
}

func AddMatrix(m1 Matrix, m2 Matrix) (Matrix, error) {
	if m1.rows() != m2.rows() || m1.cols() != m2.cols() {
		return Matrix{}, errors.New("Dimensions doesn't match")
	}

	for i, row := range m1.values {
		for j, _ := range row {
			m1.values[i][j] += m2.values[i][j]
		}
	}

	return m1, nil
}

func SubtractMatrix(m1 Matrix, m2 Matrix) (Matrix, error) {
	return AddMatrix(m1, MultiplyScalar(m2, -1))
}

func MultiplyMatrix(m1 Matrix, m2 Matrix) (Matrix, error) {
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

func Transpose(m1 Matrix) Matrix {
	result := createEmptyMatrix(m1.cols(), m1.rows())

	for i, row := range m1.values {
		for j, value := range row {
			result.values[j][i] = value
		}
	}

	return result
}

func Trace(m1 Matrix) (float64, error) {
	if m1.rows() != m1.cols() {
		return 0, errors.New("Trace is only defined for nxn square matrices")
	}

	trace := 0.0

	for i := 0; i < m1.rows(); i++ {
		trace += m1.values[i][i]
	}

	return trace, nil
}
