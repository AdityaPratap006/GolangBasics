package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 5
	fmt.Println("i is set to:", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("Th function returned:", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
