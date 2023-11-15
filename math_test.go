package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	assert.Equal(t, sum(5, 5), 10)
}

func TestSubtract(t *testing.T) {
	assert.Equal(t, subtract(10, 5), 5)
}
