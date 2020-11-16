package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidArgNr(t *testing.T) {
	var args []string

	// No arguments
	assert.Equal(t, false, isValidArgNr(args))

	// Only one argument
	args = append(args, "one")
	assert.Equal(t, true, isValidArgNr(args))

	// Two argumnets
	args = append(args, "two")
	assert.Equal(t, false, isValidArgNr(args))
}
