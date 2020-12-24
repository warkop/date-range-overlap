package daterangeoverlap

import (
	"fmt"
	"time"
)

type DateRangeOverlap struct {
}

type InputDateOverlap struct {
	StartDate *time.Time
	EndDate   *time.Time
}

func NewDateRangeOverlap() *DateRangeOverlap {
	return &DateRangeOverlap{}
}

//Validate first parameter is range date to validate
//Second parameter is allowable range
func (d *DateRangeOverlap) Validate(input, checkval InputDateOverlap) error {

	//check start date or end date not nil
	if input.StartDate == nil {
		return fmt.Errorf("Input start date cannot be null")
	}

	if input.EndDate == nil {
		return fmt.Errorf("Input end date cannot be null")
	}

	if checkval.StartDate == nil {
		return fmt.Errorf("Checkval start date cannot be null")
	}

	if checkval.StartDate == nil {
		return fmt.Errorf("Checkval start date cannot be null")
	}

	if input.StartDate.After(*checkval.EndDate) ||
		input.EndDate.After(*checkval.EndDate) ||
		input.StartDate.Before(*checkval.StartDate) ||
		input.EndDate.Before(*checkval.StartDate) {
		return fmt.Errorf("Time overlap")
	}

	return nil
}
