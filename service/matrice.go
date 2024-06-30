package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Matrice() {
	scanner := bufio.NewScanner(os.Stdin)

	var input1, input2 string

	fmt.Print("")
	if scanner.Scan() {
		input1 = scanner.Text()
	}

	fmt.Print("")
	if scanner.Scan() {
		input2 = scanner.Text()
	}

	row1 := convertToIntSlice(input1)
	row2 := convertToIntSlice(input2)

	matrix := [][]int{row1, row2}

	transposed := Trasposta(matrix)

	fmt.Println(transposed)
}

func Trasposta(m [][]int) [][]int {
	rows := len(m)
	if rows == 0 {
		return nil
	}
	cols := len(m[0])
	for i := 1; i < rows; i++ {
		if len(m[i]) != cols {
			return nil
		}
	}

	transposed := make([][]int, cols)
	for j := 0; j < cols; j++ {
		transposed[j] = make([]int, rows)
		for i := 0; i < rows; i++ {
			transposed[j][i] = m[i][j]
		}
	}
	return transposed
}

func convertToIntSlice(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		fmt.Sscanf(field, "%d", &numbers[i])
	}
	return numbers
}
