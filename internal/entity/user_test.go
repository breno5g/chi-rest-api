package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("test", "test@email.com", "senhaDaBoaESuperSegura")
	assert.Nil(t, err)

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "test", user.Username)
	assert.Equal(t, "test@email.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("test", "test@email.com", "senhaDaBoaESuperSegura")
	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("senhaDaBoaESuperSegura"))
	assert.False(t, user.ValidatePassword("senhaErrada"))
	assert.NotEqual(t, "senhaDaBoaESuperSegura", user.Password)
}
