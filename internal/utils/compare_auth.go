package utils

import (
	"crypto/sha256"
	"crypto/subtle"
)

func CompareAuth(user1, pass1, user2, pass2 string) bool {
	// Calculate SHA-256 hashes for the provided and expected
	// usernames and passwords.
	usernameHash := sha256.Sum256([]byte(user1))
	passwordHash := sha256.Sum256([]byte(pass1))
	expectedUsernameHash := sha256.Sum256([]byte(user2))
	expectedPasswordHash := sha256.Sum256([]byte(pass2))

	usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
	passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

	return usernameMatch && passwordMatch
}
