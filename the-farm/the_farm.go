package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodder FodderCalculator, numCows int) (float64, error) {
	totalFodder, err := fodder.FodderAmount(numCows)
	if err != nil {
		return 0, err
	}
	factor, err := fodder.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return totalFodder * factor / float64(numCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodder FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodder, numCows)
}

type InvalidCowsError struct {
	NumCows int
	Message string
}

func (p *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", p.NumCows, p.Message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			NumCows: numCows,
			Message: "there are no negative cows",
		}
	} else if numCows == 0 {
		return &InvalidCowsError{
			NumCows: numCows,
			Message: "no cows don't need food",
		}
	} else {
		return nil
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
