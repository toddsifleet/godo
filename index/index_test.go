package index_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/toddsifleet/godo/index"
	"github.com/toddsifleet/godo/models"
)

func getIndex() index.Index {
	idx := index.New()

	idx.AddCommand(&models.Command{
		Value:        "some-command",
		RunDirectory: "some-directory",
		Time:         time.Now(),
	})

	idx.AddCommand(&models.Command{
		Value:        "some-command",
		RunDirectory: "some-directory",
		Time:         time.Now(),
	})

	idx.AddCommand(&models.Command{
		Value:        "command-some",
		RunDirectory: "directory-some",
		Time:         time.Now(),
	})

	idx.AddCommand(&models.Command{
		Value:        "bob-barker",
		RunDirectory: "bob-barker",
		Time:         time.Now(),
	})

	idx.AddCommand(&models.Command{
		Value:        "drew-carey",
		RunDirectory: "current-directory",
		Time:         time.Now(),
	})
	return idx
}

func TestGetDirectories(t *testing.T) {
	idx := getIndex()

	result := idx.GetDirectories("current-directory", "some-directory")
	if assert.Equal(t, 1, len(result)) {
		assert.Equal(t, float64(1), result[0].Score)
		assert.Equal(t, "some-directory", result[0].Value)
		assert.Equal(t, 2, len(result[0].Times))
	}
	result = idx.GetDirectories("current-directory", "directory")
	if assert.Equal(t, 2, len(result)) {
		assert.Equal(t, "some-directory", result[0].Value)
		assert.Equal(t, "directory-some", result[1].Value)
	}
}

func TestGetCommands(t *testing.T) {
	idx := getIndex()

	result := idx.GetCommands("current-directory", "bob-barker")
	if assert.Equal(t, 1, len(result)) {
		assert.Equal(t, float64(1), result[0].Score)
		assert.Equal(t, "bob-barker", result[0].Value)
		assert.Equal(t, 1, len(result[0].Times))
	}
	result = idx.GetCommands("current-directory", "some")
	if assert.Equal(t, 2, len(result)) {
		assert.Equal(t, "some-command", result[0].Value)
		assert.Equal(t, "command-some", result[1].Value)
	}
}
