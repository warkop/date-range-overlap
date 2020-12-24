package daterangeoverlap_test

import (
	"testing"
	"time"

	"github.com/anggito12345/daterangeoverlap"
	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	time.Now()
	t.Run("Date must be oke case 1", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 20:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be oke case 2", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 06:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be oke case 3", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 03:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 07:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be oke case 4", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 18:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 22:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be overlap case 1", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 16:59")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 19:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be overlap case 2", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 06:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:01")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be overlap case 3", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 07:50")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 18:55")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Date must be overlap case 4", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 09:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 14:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Check if input start is null", func(t *testing.T) {
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 14:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: nil,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Check if input end is null", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   nil,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Check if check start is null", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 14:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: nil,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})

	t.Run("Check if check end is null", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 17:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 14:00")
		checkDate, _ := time.Parse("2006-02-01 15:04", "2020-01-01 08:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   nil,
		})

		if err != nil {
			assert.False(t, false)
		} else {
			assert.True(t, true)
		}
	})
}
