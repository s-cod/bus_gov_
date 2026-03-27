package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func tt() []int64 {

	ar := [10]int64{}
	sl := ar[5:]

	return sl

}

func TestTT(t *testing.T) {
	r := tt()
	result := []int64{0, 0, 0, 0, 0}
	assert.Equal(t, r, result)

}
