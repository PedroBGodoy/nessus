package user

import "log"

// User define user structure
type User struct {
	UserID   string
	Name     string
	Email    string
	Password string
}

// CreateUser create a new user
func CreateUser(name string, email string, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user, nil
}

// LoginUser login user and returns JWT
func LoginUser(email string, password string) (string, error) {
	token, err := generateToken(email)
	if err != nil {
		log.Fatal("error when generating token: ", err)
	}

	return token, nil
}

// Authenticate validates user token
func Authenticate(token string) (*User, error) {
	err := validatesToken(token)
	if err != nil {
		return nil, err
	}

	u := &User{
		UserID: "123",
		Name:   "teste2",
		Email:  "teste2",
	}

	return u, nil
}
