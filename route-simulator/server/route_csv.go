package main

import (
	"encoding/csv"
	"io"
	"strconv"
)

type csvReader struct{}

func (r *csvReader) Read(reader io.Reader) ([]Point, error) {
	// TODO: implement route csv reader
	var points []Point
	lines, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		lng, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			return nil, err
		}
		lat, err := strconv.ParseFloat(line[1], 64)
		if err != nil {
			return nil, err
		}
		point := Point{lng, lat}
		points = append(points, point)
	}
	return points, nil
}

func createCSVReader() Reader {
	return &csvReader{}
}

func init() {
	RegisterReader("csv", createCSVReader)
}
