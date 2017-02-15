package csvlib

import (
	"encoding/csv"
	"io"
	"os"
)

// ReadFiles reads files []string and outputs 3D slice of all
// the data contained in the csv files. 3D slice should be ready
// to manipulate in any way your heart desires.
func ReadFiles(files []string) ([][][]string, error) {
	var list [][][]string
	for _, file := range files {
		csvFile, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer csvFile.Close()

		csvCnts, err := parseCSV(csvFile)
		if err != nil {
			return nil, err
		}
		list = append(list, csvCnts)
	}
	return list, nil
}

func parseCSV(csvFile io.Reader) ([][]string, error) {
	var csvCnts [][]string
	f := csv.NewReader(csvFile)

	for {
		record, err := f.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		csvCnts = append(csvCnts, record)
	}
	return csvCnts, nil
}
