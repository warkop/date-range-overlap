package daterangeoverlap

import (
	"fmt"
	"time"
)

// DateRangeOverlap is constructor
type DateRangeOverlap struct {
}

// InputDateOverlap is struct for save start and end date
type InputDateOverlap struct {
	StartDate *time.Time
	EndDate   *time.Time
}

// NewDateRangeOverlap initial action
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

	if checkval.EndDate == nil {
		return fmt.Errorf("Checkval end date cannot be null")
	}

	// this is a main logic, would check input date is overlap with existing data or not
	if input.StartDate.Before(*checkval.EndDate) && input.EndDate.After(*checkval.StartDate) {
		return fmt.Errorf("Time overlap")
	}

	return nil
}
