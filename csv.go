package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines the generic statistical function
type statsFunc func([]float64) float64

func sum(data []float64) float64 {
	sum := 0.0

	for _, value := range data {
		sum += value
	}

	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV Reader used to read in data from CSV files
	cr := csv.NewReader(r)
	cr.ReuseRecord = true

	// Adjusting for 0 based index
	column--

	var data []float64

	// Looping through all records
	for i := 0; ; i++ {
		row, err := cr.Read()
		if err == io.EOF {
			// End of file
			break
		} else if err != nil {
			// Error reading CSV file
			return nil, fmt.Errorf("Cannot read CSV data: %w", err)
		}

		if i == 0 {
			// Skip header row
			continue
		}

		// Checking number of columns in CSV file
		if len(row) <= column {
			// File does not have that many columns
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}

		// Try to convert data read into a float number
		value, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}

		data = append(data, value)
	}

	// Return the slice fo float64 and nil error
	return data, nil
}
