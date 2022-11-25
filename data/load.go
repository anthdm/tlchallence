package data

import (
	"encoding/gob"
	"log"
	"os"
)

func LoadOrdersFromFile(filepath string) map[float64]float64 {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data map[float64]float64
	if err := gob.NewDecoder(f).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return data
}
