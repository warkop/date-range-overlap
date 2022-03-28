package daterangeoverlap

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	HourMinute = "15:04"
)

func TestFirst(t *testing.T) {
	time.Now()

	t.Run("Date must be oke case 1", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "17:00")
		startEndDate, _ := time.Parse(HourMinute, "20:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be oke case 2", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "06:00")
		startEndDate, _ := time.Parse(HourMinute, "08:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be oke case 3", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "03:00")
		startEndDate, _ := time.Parse(HourMinute, "07:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be oke case 4", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "18:00")
		startEndDate, _ := time.Parse(HourMinute, "22:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be overlap case 1", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "16:59")
		startEndDate, _ := time.Parse(HourMinute, "19:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be overlap case 2", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "06:00")
		startEndDate, _ := time.Parse(HourMinute, "08:01")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be overlap case 3", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "07:50")
		startEndDate, _ := time.Parse(HourMinute, "18:55")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Date must be overlap case 4", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "09:00")
		startEndDate, _ := time.Parse(HourMinute, "14:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Check if input start is null", func(t *testing.T) {
		startEndDate, _ := time.Parse(HourMinute, "14:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: nil,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Check if input end is null", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "17:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   nil,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Check if check start is null", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "17:00")
		startEndDate, _ := time.Parse(HourMinute, "14:00")
		checkEndDate, _ := time.Parse(HourMinute, "17:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: nil,
			EndDate:   &checkEndDate,
		})

		assert.NoError(t, err)
	})

	t.Run("Check if check end is null", func(t *testing.T) {
		startDate, _ := time.Parse(HourMinute, "17:00")
		startEndDate, _ := time.Parse(HourMinute, "14:00")
		checkDate, _ := time.Parse(HourMinute, "08:00")
		err := NewDateRangeOverlap().Validate(InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   nil,
		})

		assert.NoError(t, err)
	})
}
