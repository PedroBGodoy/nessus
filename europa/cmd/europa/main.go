package main

import (
	"log"

	"github.com/nessus/europa/internal/user"
)

func main() {
	// u, _ := user.CreateUser("teste", "teste", "teste")

	// log.Printf("Name: %s, Email: %s, Password: %s", u.Name, u.Email, u.Password)

	token, _ := user.LoginUser("teste", "teste")
	log.Printf("Token: %s", token)

	err := user.Authenticate(token)
	if err != nil {
		log.Fatalf("User not authenticated: %s", err)
	}

	log.Print("User autenticated!")
}
