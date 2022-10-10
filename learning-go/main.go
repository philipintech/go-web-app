package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
	// note: capitalizing varibable/function allows it to be visible to outside of the "main" package
}

func main() {
	user := User{
		FirstName:   "Saul",
		LastName:    "Goodman",
		PhoneNumber: "416-614-4616",
	}
	log.Println(user.FirstName, user.LastName, "BirthDate", user.BirthDate)
}

// func saySomething(s3 string) (string, string) {
// 	log.Println("s from saySomething func is", s)
// 	return s3, "world"
// }
