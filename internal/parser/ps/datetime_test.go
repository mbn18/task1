package ps

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvertLstart(t *testing.T) {
	items := []struct {
		DateString string
		DateTime   time.Time
	}{
		{
			DateString: "Thu Feb 27 05:07:12 2023",
			DateTime:   time.Date(2023, time.February, 27, 5, 7, 12, 0, time.UTC),
		},
		{
			DateString: "Mon Jun 15 14:32:45 2024",
			DateTime:   time.Date(2024, time.June, 15, 14, 32, 45, 0, time.UTC),
		},
		{
			DateString: "Wed Nov 12 09:28:33 2025",
			DateTime:   time.Date(2025, time.November, 12, 9, 28, 33, 0, time.UTC),
		},
	}

	for _, item := range items {
		s, err := convertLstart(item.DateString)
		assert.NoError(t, err)
		assert.Equal(t, item.DateTime, s)
	}

	_, err := convertLstart("jibrish")
	assert.Error(t, err)
}

func TestConvertCpuTime(t *testing.T) {
	items := []struct {
		str      string
		duration time.Duration
	}{
		{str: "00:00:01", duration: time.Second},
		{str: "01:01:01", duration: time.Hour + time.Minute + time.Second},
		{str: "01:02:03", duration: time.Hour + time.Minute*2 + time.Second*3},
	}
	for _, item := range items {
		s, err := convertCpuTime(item.str)
		assert.NoError(t, err)
		assert.Equal(t, item.duration, s)
	}

	_, err := convertCpuTime("00:21")
	assert.Error(t, err)
}
