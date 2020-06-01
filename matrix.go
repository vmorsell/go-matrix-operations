package matrix

import "errors"

type Matrix struct {
	values [][]float64
}

func (m1 Matrix) rows() int {
	return len(m1.values)
}

func (m1 Matrix) cols() int {
	return len(m1.values[0])
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
