package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInstance(t *testing.T) {
	instanceA := NewInstance()
	instanceA["name"] = "duc"

	instanceB := NewInstance()

	assert.Equal(t, instanceA, instanceB)
	assert.Equal(t, "duc", instanceB["name"])
}
