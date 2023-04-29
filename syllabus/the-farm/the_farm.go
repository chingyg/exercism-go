package thefarm

import (
	"errors"
	"fmt"
)

// TODO: Define the SillyNephewError type here.

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {

	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	fodder, err := weightFodder.FodderAmount()

	// Scale is broken
	if err == ErrScaleMalfunction {
		if fodder < 0 {
			return 0, errors.New("negative fodder")
		} else {
			// Double the fodder
			fodder *= 2
		}
	} else if err != nil {
		// Generic error
		return 0, err
	}

	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0, &SillyNephewError{cows}
	}

	return fodder / float64(cows), nil

}
