package option

import (
	"design_patterns/option/option"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T)  {
	options := option.New(option.SetAge(21), option.SetId(1), option.SetName("duc"))
	assert.Equal(t, int64(21), options.Age)
	assert.Equal(t, "duc", options.Name)
	assert.Equal(t, 1, options.Id)
}