package security

import "golang.org/x/crypto/bcrypt"

// Hash encrypt password provided
func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

// VerifyPassword compare hashed password agains plain password
func VerifyPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
