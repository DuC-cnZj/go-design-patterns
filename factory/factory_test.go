package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConst(t *testing.T) {
	assert.Equal(t, 1, int(AppleFactory))
	assert.Equal(t, 2, int(BananaFactory))
	assert.Equal(t, 4, int(PeachFactory))
}

func TestCreate(t *testing.T)  {
	apple := Create(AppleFactory)
	assert.Equal(t, "apple", Get())
}
