package matrix

import (
	"fmt"
	"math"
)

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

type Matrix [][]float64

func (m Matrix) Print() {
	for _, row := range m {
		for _, el := range row {
			fmt.Printf("%f ", el)
		}
		fmt.Println("")
	}
}

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Columns() int {
	return len(m[0])
}

func (m Matrix) ExcludeColumn(column int) (Matrix, error) {

	if !InBetween(column, 1, m.Columns()) {
		return Matrix{}, fmt.Errorf("input not in range")
	}

	result := make(Matrix, m.Rows())
	for i, row := range m {
		for j, el := range row {
			if j == column-1 {
				continue
			}
			result[i] = append(result[i], el)
		}
	}
	return result, nil
}

func (m Matrix) ExcludeRow(row int) (Matrix, error) {
	if !InBetween(row, 1, m.Rows()) {
		return Matrix{}, fmt.Errorf("input not in range")
	}

	var result Matrix
	for i, r := range m {
		if i == row-1 {
			continue
		}
		result = append(result, r)
	}
	return result, nil
}

func (m Matrix) SubMatrix(column_start int, column_end int, row_start int, row_end int) (Matrix, error) {

	if !InBetween(column_start, 1, column_end) ||
		!InBetween(column_end, column_start, m.Columns()) ||
		!InBetween(row_start, 1, row_end) ||
		!InBetween(row_end, row_start, m.Rows()) {
		return Matrix{}, fmt.Errorf("inputs not in range")
	}

	partial_matrix := m[row_start+1 : row_end+1]
	for i, row := range partial_matrix {
		partial_matrix[i] = row[column_start+1 : column_end+1]
	}

	return partial_matrix, nil
}

func (m Matrix) IsMatrix() bool {

	if m.Rows() == 0 {
		return false
	}

	for _, row := range m {

		if len(m[0]) != len(row) {
			return false
		}

	}

	return true
}

func (m Matrix) IsSquare() bool {
	return m.Columns() == m.Rows()
}

func (m Matrix) Det() (float64, error) {

	if !m.IsMatrix() || !m.IsSquare() {
		return -1, fmt.Errorf("determinant is not defined for the input [Matrix: %t][Square: %t]",
			m.IsMatrix(), m.IsSquare())
	}

	if m.Rows() == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0], nil
	}

	// if rows are more than 2
	partial_matrix, err := m.ExcludeRow(1)
	if err != nil {
		return -1, err
	}

	var temp float64 = 0
	for i, el := range m[0] {

		reduced_matrix, err := partial_matrix.ExcludeColumn(i + 1)
		if err != nil {
			return -1, err
		}

		partial_det, err := reduced_matrix.Det()
		if err != nil {
			return -1, err
		}
		temp = temp + partial_det*el*math.Pow(-1, float64(i))
	}

	return temp, nil
}
