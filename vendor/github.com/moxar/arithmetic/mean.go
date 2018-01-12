package arithmetic

import (
	"errors"
	"fmt"
)

func init() {
	RegisterFunction("mean", mean)
}

// mean is a function that returns the mean of the provided inputs.
func mean(args ...interface{}) (interface{}, error) {

	if len(args) == 0 {
		return nil, errors.New("mean requires at least one argument")
	}

	// Ensure each argument is a float, or a "variable" float.
	var sum float64
	for _, a := range args {
		switch t := a.(type) {
		case float64:
			sum += t
		case variable:
			v, ok := t.value.(float64)
			if !ok {
				return nil, fmt.Errorf("mean requires numeric arguments, %s given", t)
			}
			sum += v
		default:
			return nil, fmt.Errorf("mean requires numeric arguments, %v given", a)
		}
	}

	return sum / float64(len(args)), nil
}
