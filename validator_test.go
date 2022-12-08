package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePasswrd(t *testing.T) {

	t.Run("empty", func(t *testing.T) {
		result := ValidatePassword("", 8)
		assert.False(t, result)
	})

	t.Run("all characters", func(t *testing.T) {
		result := ValidatePassword("Aa4[asdf", 8)
		assert.True(t, result)
	})

	t.Run("no specials", func(t *testing.T) {
		result := ValidatePassword("Aa45asdf", 8)
		assert.False(t, result)

	})

	t.Run("no upper case", func(t *testing.T) {
		result := ValidatePassword("[a45asdf", 8)
		assert.False(t, result)
	})

	t.Run("specific case", func(t *testing.T) {
		result := ValidatePassword("H^9AdVLjTEF5Tm", 8)
		assert.True(t, result)
	})
}
