package ps_test

import (
	"bytes"
	"github.com/mbn18/dream/internal/entity"
	"github.com/mbn18/dream/internal/parser/ps"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	header           = []byte("USER         PID %CPU %MEM    VSZ   RSS TTY      STAT  STARTED     TIME COMMAND")
	headerMissingCol = []byte("USER         PID %CPU %MEM    VSZ   RSS TTY      STAT STARTED   TIME")
	headerTypo       = []byte("USEX         PID %CPU %MEM    VSZ   RSS TTY      STAT STARTED   TIME COMMAND")

	line1     = []byte("root           1  1.1  2.1  65452 30880 ?        Ss   Tue Feb 11 15:36:39 2025 00:00:19 cmd1 -arg1 -arg2")
	line1Date = time.Date(2025, 2, 11, 15, 36, 39, 0, time.UTC)

	line2     = []byte("user1           2  0.5  0.6      0     0 ?        S    Wed Jan 26 13:50:28 2024   01:00:02 cmd2")
	line2Date = time.Date(2024, 1, 26, 13, 50, 28, 0, time.UTC)

	lineMissingField = []byte("root           2  0.0  0.0      0     0 ?        S    0:00 [kthreadd]")
)

func TestParse(t *testing.T) {
	input := concatWithNewLine(header, line1, line2)

	output, err := ps.Parse(input)
	assert.NoError(t, err)
	assert.Len(t, output, 2)

	p1 := &entity.Process{
		User:    "root",
		PID:     1,
		CPU:     1.1,
		Memory:  2.1,
		VSZ:     65452,
		RSS:     30880,
		TTY:     "?",
		Stat:    "Ss",
		Start:   line1Date,
		CPUTime: time.Second * 19,
		Command: "cmd1",
		Args:    "-arg1 -arg2",
	}
	assert.Equal(t, p1, output[0])

	p2 := &entity.Process{
		User:    "user1",
		PID:     2,
		CPU:     0.5,
		Memory:  0.6,
		VSZ:     0,
		RSS:     0,
		TTY:     "?",
		Stat:    "S",
		Start:   line2Date,
		CPUTime: time.Hour + time.Second*2,
		Command: "cmd2",
	}
	assert.Equal(t, p2, output[1])
}

func TestParse_missingFields(t *testing.T) {
	input := concatWithNewLine(header, line1, lineMissingField)

	_, err := ps.Parse(input)
	assert.ErrorContains(t, err, "missing field")

}

func TestParse_EmptyInput(t *testing.T) {
	_, err := ps.Parse([]byte{})
	assert.ErrorContains(t, err, "no content")
}

func TestParse_OnlyHeader(t *testing.T) {
	_, err := ps.Parse(header)
	assert.ErrorContains(t, err, "no content")
}

func TestParse_InvalidHeader(t *testing.T) {
	input := concatWithNewLine(headerMissingCol, line1, line2)
	_, err := ps.Parse(input)
	assert.ErrorContains(t, err, "header format is invalid")

	input = concatWithNewLine(headerTypo, line1, line2)
	_, err = ps.Parse(input)
	assert.ErrorContains(t, err, "header format is invalid")
}

func concatWithNewLine(lines ...[]byte) []byte {
	buffer := bytes.Buffer{}
	for _, line := range lines {
		buffer.Write(line)
		buffer.WriteByte('\n')
	}
	return buffer.Bytes()
}
