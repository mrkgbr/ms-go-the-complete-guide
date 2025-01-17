package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, stringValue := range strings {
		floatPrice, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("an error occurred during convert: failed to convert string to float")
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
