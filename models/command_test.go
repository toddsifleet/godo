package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandFromLogLineInvalid(t *testing.T) {
	_, err := CommandFromLogLine("adsf")
	assert.EqualError(t, err, "invalid line")
}

func TestCommandFromLogLineInvalidTimeStamp(t *testing.T) {
	_, err := CommandFromLogLine("un Apr  8 17:49:42 PDT 2018__GODO_SPLIT__FOOBAR__GODO_SPLIT__FOOBAR")
	assert.EqualError(
		t,
		err,
		`parsing time "un Apr  8 17:49:42 PDT 2018_" as "Mon Jan _2 15:04:05 MST 2006": cannot parse "un Apr  8 17:49:42 PDT 2018_" as "Mon"`,
	)
}

func TestCommandFromLogLineValid(t *testing.T) {
	result, err := CommandFromLogLine("Sun Apr  8 17:49:42 PDT 2018__GODO_SPLIT__RUN_DIRECTORY__GODO_SPLIT__COMMAND")
	if assert.NoError(t, err, "should not error") {
		assert.Equal(t, "RUN_DIRECTORY", result.RunDirectory)
		assert.Equal(t, "COMMAND", result.Value)
	}
}
