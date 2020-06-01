package matrix

import "errors"

type Matrix struct {
	Values [][]float64
}

func AddScalar(m1 Matrix, scalar float64) Matrix {
	result := Matrix{[][]float64{}}

	for i, rows := range m1.Values {
		values := []float64{}
		for j, _ := range rows {
			values = append(values, m1.Values[i][j]+scalar)
		}
		result.Values = append(result.Values, values)
	}

	return result
}

func SubtractScalar(m1 Matrix, scalar float64) Matrix {
	return AddScalar(m1, -scalar)
}

func MultiplyScalar(m1 Matrix, scalar float64) Matrix {
	result := Matrix{[][]float64{}}

	for i, row := range m1.Values {
		values := []float64{}
		for j, _ := range row {
			values = append(values, m1.Values[i][j]*scalar)
		}
		result.Values = append(result.Values, values)
	}

	return result
}
