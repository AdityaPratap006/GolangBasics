package main

import "log"

func main() {

	// demoForLoop()

	// demoRangingOverSlice()

	// demoRangingOverMap()

	// demoRangingOverString()

	demoRangingOverCustomObjects()
}

func demoForLoop() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func demoRangingOverSlice() {
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for _, animal := range animals {
		log.Println(animal)
	}
}

func demoRangingOverMap() {
	animals := make(map[string]string)

	animals["dog"] = "Zig"
	animals["cat"] = "Fluffy"

	for animalType, animalName := range animals {
		log.Println(animalType, animalName)
	}
}

func demoRangingOverString() {
	var firstLine string = "Once upon a midnight dreary"

	for i, letter := range firstLine {
		log.Println(i, ":", letter)
	}

}

func demoRangingOverCustomObjects() {

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Wick", "john.wick@x.com", 40})
	users = append(users, User{"Mary", "Jones", "marry.jones@x.com", 25})
	users = append(users, User{"Sally", "Brown", "sally.brown@x.com", 30})
	users = append(users, User{"Alex", "Anderson", "alex.anderson@x.com", 22})

	for _, u := range users {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}

}
