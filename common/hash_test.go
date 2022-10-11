package common

import (
	"e-wallet-app/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := util.RandomString(12)
	wrongPassword := util.RandomString(10)
	tests := []struct {
		name                string
		password            string
		expectHashPassword  func(t *testing.T, hash string, err error)
		expectCheckPassword func(t *testing.T, err error)
	}{
		{
			name:     "correct password",
			password: password,
			expectHashPassword: func(t *testing.T, hash string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hash)
			},
			expectCheckPassword: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
		{
			name:     "wrong password",
			password: wrongPassword,
			expectHashPassword: func(t *testing.T, hash string, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hash)
			},
			expectCheckPassword: func(t *testing.T, err error) {
				require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
			},
		},
	}

	var (
		hash string
		err  error
	)

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			hash, err = HashPassword(password)
			test.expectHashPassword(t, hash, err)

			err = CheckPassword(test.password, hash)
			test.expectCheckPassword(t, err)
		})
	}
}
