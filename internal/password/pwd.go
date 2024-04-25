package pwd

import (
	"github.com/go-crypt/crypt"
	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
)

func ValidatePassword(hash, password string) bool {
	var (
		valid bool
		err   error
	)

	if valid, err = crypt.CheckPassword(password, hash); err != nil {
		return valid
	}

	return valid
}

func Encode(psw string) (hash string) {
	var (
		hasher *argon2.Hasher
		err    error
		digest algorithm.Digest
	)

	if hasher, err = argon2.New(
		argon2.WithProfileRFC9106LowMemory(),
	); err != nil {
		panic(err)
	}

	if digest, err = hasher.Hash(psw); err != nil {
		panic(err)
	}

	return digest.Encode()
}
