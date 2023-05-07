// https://www.hackerrank.com/challenges/encryption/problem
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func findMinMatrixSize(floor, ceil, l int) (int, int) {
	for i := floor; i <= ceil; i++ {
		for j := i; j <= ceil; j++ {
			if i*j >= l {
				return i, j
			}
		}
	}

	return ceil, ceil
}

func buildMatrix(nRow, nCol int, s string) []string {
	matrix := make([]string, nRow)

	for i := range s {
		insertingRow := int(i / nCol)
		if insertingRow < nRow-1 {
			matrix[insertingRow] = s[insertingRow*nCol : (insertingRow+1)*nCol]
		} else {
			matrix[insertingRow] = s[insertingRow*nCol:]
		}
	}

	return matrix
}

func encryptMatrix(matrix []string, nRow, nCol int) string {
	var resByteBuf bytes.Buffer

	lastStrLen := len(matrix[nRow-1])
	for i := 0; i < nCol; i++ {
		rowIndUpperBound := nRow
		if lastStrLen-1 < i {
			rowIndUpperBound = nRow - 1
		}

		for j := 0; j < rowIndUpperBound; j++ {
			// fmt.Printf("col %d and row %d is : %s\n",i,  j, matrix[j])
			resByteBuf.WriteByte(matrix[j][i])
		}
		if i < nCol-1 {
			resByteBuf.WriteString(" ")
		}
	}

	return resByteBuf.String()
}

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	sqrLen := math.Sqrt(float64(len(s)))

	fSqrLen := int(math.Floor(sqrLen))
	cSqrLen := int(math.Ceil(sqrLen))

	nRow, nCol := findMinMatrixSize(fSqrLen, cSqrLen, len(s))

	matrix := buildMatrix(nRow, nCol, s)

	s = encryptMatrix(matrix, nRow, nCol)

	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
