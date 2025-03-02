package ps

import (
	"bytes"
	"errors"
	"github.com/mbn18/dream/internal/entity"
	"strconv"
)

const expectedFields = 11

// To normalize the output on both Linux/Mac of:
// ps -e -o user,pid,%cpu,%mem,vsz,rss,tty=TTY,stat,start,time,args=COMMAND|head -n 1
var headerColList = bytes.Fields([]byte("USER PID %CPU %MEM VSZ RSS TTY STAT STARTED TIME COMMAND"))

func Parse(input []byte) ([]*entity.Process, error) {

	lines := bytes.Split(input, []byte("\n"))
	if len(lines) < 2 {
		return nil, errors.New("no content")
	}

	if err := validateHeader(lines[0]); err != nil {
		return nil, err
	}

	// Create a list with predefined length, ommit header and last new line
	list := make([]*entity.Process, len(lines)-2)

	var err error
	for num, line := range lines[1 : len(lines)-1] {
		if list[num], err = parseLine(line); err != nil {
			return nil, err
		}
	}
	return list, nil
}

func validateHeader(header []byte) error {
	fields := bytes.Fields(header)
	if len(fields) != len(headerColList) {
		return errors.New("header format is invalid")
	}
	for n, field := range fields {
		if !bytes.Equal(field, headerColList[n]) {
			return errors.New("header format is invalid")

		}
	}
	return nil
}

func parseLine(line []byte) (*entity.Process, error) {
	fields := bytes.Fields(line)

	if len(fields) < expectedFields {
		return nil, errors.New("missing fields")
	}

	p := &entity.Process{
		User:    string(fields[0]),
		TTY:     string(fields[6]),
		Stat:    string(fields[7]),
		Command: string(fields[14]),
		Args:    concatWithSpace(fields[15:]),
	}

	var err error
	if p.PID, err = strconv.Atoi(string(fields[1])); err != nil {
		return nil, errors.New("invalid PID value")
	}
	if p.CPU, err = strconv.ParseFloat(string(fields[2]), 64); err != nil {
		return nil, errors.New("invalid CPU value")
	}
	if p.Memory, err = strconv.ParseFloat(string(fields[3]), 64); err != nil {
		return nil, errors.New("invalid MEM value")
	}
	if p.VSZ, err = strconv.Atoi(string(fields[4])); err != nil {
		return nil, errors.New("invalid VSZ value")
	}
	if p.RSS, err = strconv.Atoi(string(fields[5])); err != nil {
		return nil, errors.New("invalid RSS value")
	}
	if p.Start, err = convertLstart(concatWithSpace(fields[8:13])); err != nil {
		return nil, errors.New("invalid STARTED value")
	}
	if p.CPUTime, err = convertCpuTime(string(fields[13])); err != nil {
		return nil, err
	}
	return p, nil
}

func concatWithSpace(list [][]byte) string {
	return string(bytes.Join(list, []byte(" ")))
}
