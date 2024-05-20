package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var sliceOfFloats []float64
	for _, lineVal := range strings {
		floatPrice, err := strconv.ParseFloat(lineVal, 64)

		if err != nil {
			return nil, errors.New("error")
		}

		sliceOfFloats = append(sliceOfFloats, floatPrice)
	}
	return sliceOfFloats, nil
}
