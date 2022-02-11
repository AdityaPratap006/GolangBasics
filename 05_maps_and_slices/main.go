package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	demoMap()

	demoSlice()

	demoSlice2()
}

func demoMap() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "John",
		LastName:  "Wick",
	}

	myMap["hero"] = me

	log.Println(myMap["hero"])
}

func demoSlice() {
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)

	sort.Ints(mySlice)

	log.Println(mySlice)

}

func demoSlice2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	log.Println(numbers)
	log.Println(numbers[0:2])
}
