package main

import (
	"bufio"
	"encoding/csv"
	"gonum.org/v1/gonum/mat"
	"io"
	"os"
	"strconv"
)

func LoadIrisData(filepath string) (X *mat.Dense, Y *mat.Dense) {
	if csvFile, fileErr := os.Open(filepath); fileErr == nil {
		reader := csv.NewReader(bufio.NewReader(csvFile))

		xData := []float64{}
		yData := []int{}
		for {
			if dataRow, readerErr := reader.Read(); readerErr != nil && readerErr == io.EOF {
				break
			} else {
				for i := 0; i < len(dataRow)-1; i += 1 {
					if el, parseErr := strconv.ParseFloat(dataRow[i], 64); parseErr == nil {
						xData = append(xData, el)
					}
				}

				if el, parseErr := strconv.ParseInt(dataRow[len(dataRow)-1], 10, 64); parseErr == nil {
					yData = append(yData, int(el))
				}
			}

		}

		inputDim := 4
		N := len(yData)
		X = mat.NewDense(N, inputDim, xData)

		numClass := 3
		Y = mat.NewDense(N, numClass, nil)
		for i := 0; i < N; i += 1 {
			Y.Set(i, yData[i], 1)
		}

		return X, Y
	}

	return nil, nil
}
