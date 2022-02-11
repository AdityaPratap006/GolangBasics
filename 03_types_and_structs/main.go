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
}

func main() {
	user := User{
		FirstName:   "John",
		LastName:    "Wick",
		PhoneNumber: "1 555 555-1212",
		BirthDate:   time.Date(1970, time.January, 6, 19, 10, 15, 120, &time.Location{}),
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}

func Whatever() {

}
