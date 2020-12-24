package daterangeoverlap_test

import (
	"testing"
	"time"

	"github.com/anggito12345/daterangeoverlap"
)

func TestFirst(t *testing.T) {
	time.Now()
	t.Run("Date must be overlap", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:00:01")
		checkDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:00:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil && err.Error() != "Time Overlap" {
			t.Log("OK")
		} else {
			t.Error("Should be error")
		}
	})

	t.Run("Date must be oke", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 16:00:00")
		checkDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:00:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil {
			t.Error(err)
		} else {
			t.Log("OK!")
		}
	})

	t.Run("Date must be overlap case 2", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 07:59:59")
		startEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 16:00:00")
		checkDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:00:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil && err.Error() != "Time Overlap" {
			t.Log("OK")
		} else {
			t.Error("Should be error")
		}
	})

	t.Run("Date must be overlap case 3", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 07:50:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 07:55:00")
		checkDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:00:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil && err.Error() != "Time Overlap" {
			t.Log("OK")
		} else {
			t.Error("Should be error")
		}
	})

	t.Run("Date must be overlap case 4", func(t *testing.T) {
		startDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:01:00")
		startEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:02:00")
		checkDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 08:00:00")
		checkEndDate, _ := time.Parse("2006-02-01 15:04:05", "2020-01-01 17:00:00")
		err := daterangeoverlap.NewDateRangeOverlap().Validate(daterangeoverlap.InputDateOverlap{
			StartDate: &startDate,
			EndDate:   &startEndDate,
		}, daterangeoverlap.InputDateOverlap{
			StartDate: &checkDate,
			EndDate:   &checkEndDate,
		})

		if err != nil && err.Error() != "Time Overlap" {
			t.Log("OK")
		} else {
			t.Error("Should be error")
		}
	})
}
