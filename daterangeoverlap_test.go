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
	tableTest := []struct {
		name           string
		inputStartTime string
		inputEndTime   string
		checkStartTime string
		checkEndTime   string
		wantError      bool
	}{
		{
			name:           "Date must be oke case 1",
			inputStartTime: "17:00",
			inputEndTime:   "20:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      false,
		},
		{
			name:           "Date must be oke case 2",
			inputStartTime: "03:00",
			inputEndTime:   "07:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      false,
		},
		{
			name:           "Date must be oke case 2",
			inputStartTime: "06:00",
			inputEndTime:   "08:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      false,
		},
		{
			name:           "Date must be oke case 3",
			inputStartTime: "03:00",
			inputEndTime:   "07:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      false,
		},
		{
			name:           "Date must be oke case 4",
			inputStartTime: "18:00",
			inputEndTime:   "22:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      false,
		},
		{
			name:           "Date must be overlap case 1",
			inputStartTime: "16:59",
			inputEndTime:   "19:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Date must be overlap case 2",
			inputStartTime: "06:00",
			inputEndTime:   "08:01",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Date must be overlap case 3",
			inputStartTime: "07:50",
			inputEndTime:   "18:55",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Date must be overlap case 4",
			inputStartTime: "09:00",
			inputEndTime:   "14:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Check if input start is null",
			inputStartTime: "",
			inputEndTime:   "14:00",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Check if input end is null",
			inputStartTime: "17:00",
			inputEndTime:   "",
			checkStartTime: "08:00",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Check if check start is null",
			inputStartTime: "17:00",
			inputEndTime:   "14:00",
			checkStartTime: "",
			checkEndTime:   "17:00",
			wantError:      true,
		},
		{
			name:           "Check if check end is null",
			inputStartTime: "14:00",
			inputEndTime:   "17:00",
			checkStartTime: "08:00",
			checkEndTime:   "",
			wantError:      true,
		},
	}

	for _, tt := range tableTest {
		t.Run(tt.name, func(t *testing.T) {
			inputStartDate, _ := time.Parse(HourMinute, tt.inputStartTime)
			inputEndDate, _ := time.Parse(HourMinute, tt.inputEndTime)
			checkStartDate, _ := time.Parse(HourMinute, tt.checkStartTime)
			checkEndDate, _ := time.Parse(HourMinute, tt.checkEndTime)
			err := NewDateRangeOverlap().Validate(InputDateOverlap{
				StartDate: &inputStartDate,
				EndDate:   &inputEndDate,
			}, InputDateOverlap{
				StartDate: &checkStartDate,
				EndDate:   &checkEndDate,
			})

			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
