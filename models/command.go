package models

import (
	"errors"
	"strings"
	"time"
)

const TIMESTAMP_LENGTH = len(time.UnixDate)

type Command struct {
	Time         time.Time
	RunDirectory string
	Value        string
}

func parseCommand(command []string) string {
	return strings.Trim(strings.Join(command, " "), " ")
}

func CommandFromLogLine(line string) (*Command, error) {
	if len(line) < TIMESTAMP_LENGTH+5 {
		return nil, errors.New("invalid line")
	}
	t, err := time.Parse(time.UnixDate, line[0:TIMESTAMP_LENGTH])
	if err != nil {
		return nil, err
	}

	splits := strings.Split(line, "__GODO_SPLIT__")
	if len(splits) < 3 {
		return nil, errors.New("invalid line")
	}
	return &Command{
		Time:         t,
		Value:        parseCommand(splits[2:]),
		RunDirectory: splits[1],
	}, nil
}
