package main

import (
	"github.com/AdityaPratap006/myniceprogram/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType

	myVar.TypeName = "Some Name"

	log.Println(myVar)
}
