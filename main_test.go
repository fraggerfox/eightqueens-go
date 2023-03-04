package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	expectedExit := 0
	actualExit := realMain(&buf)

	expectedMessage := "Found 92 solutions.\n"
	actualMessage := buf.String()

	assert.Equal(t, expectedExit, actualExit)
	assert.Equal(t, expectedMessage, actualMessage)
}
