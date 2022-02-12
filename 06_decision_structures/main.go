package main

import "log"

func main() {
	// demo1()

	// demo2()

	demo3()
}

func demo1() {
	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is NOT cat")
	}
}

func demo2() {
	myNum := 101
	isTrue := true

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to false")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum == 101 || isTrue {
		log.Println("2")
	} else if myNum > 100 && !isTrue {
		log.Println("3")
	}
}

func demo3() {

	myVar := "dog"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")

	case "dog":
		log.Println("cat is set to dog")

	case "fish":
		log.Println("cat is set to fish")

	default:
		log.Println("cat is something else")
	}
}
