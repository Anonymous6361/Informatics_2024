package laba8

import (
	"io"
	"os"
	"strconv"
	"strings"

	"isuct.ru/informatics2022/lab4"
)

const PathTask1 = "lab8/input.txt"

func task1() error {
	file, err := os.Open(PathTask1)
	if err != nil {
		return err
	}
	defer file.Close()

	var text string
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		text = string(data[:n])
	}

	var result []float64
	Parameters := strings.Split(text, "\r\n")
	for _, P := range Parameters {
		number, err := strconv.ParseFloat(P, 64)
		if err != nil {
			return err
		}
		result = append(result, number)
	}

	xMin:= result[0]
	xMax := result[1]
	xDelta := result[2]
	laba4.TaskA(xMin, xMax, xDelta)
	laba4.TaskB(result[3:7])

	return nil