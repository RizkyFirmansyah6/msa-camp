package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type jsonReader struct{}

func (r *jsonReader) Read(reader io.Reader) ([]Point, error) {
	var points []Point
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(reader)
	// unmarshal byteArray which contains
	// jsonFile's content into 'points'
	err := json.Unmarshal(byteValue, &points)
	if err != nil {
		return nil, err
	}
	return points, nil
}

func createJSONReader() Reader {
	return &jsonReader{}
}

func init() {
	RegisterReader("json", createJSONReader)
}
