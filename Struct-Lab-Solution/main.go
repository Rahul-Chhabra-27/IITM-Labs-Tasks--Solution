package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	firstNames := [5]string{"John", "Mark", "Tom", "Ben", "stuart"}
	lastNames := [5]string{"Doe", "Wood", "Curran", "Foakes", "Broad"}
	ages := [5]int{12, 23, 43, 18, 14}

	var personsArray [5]Person

	for index := 0; index < 5; index++ {
		person := Person{firstName: firstNames[index],
			lastName: lastNames[index], age: ages[index]}
		personsArray[index] = person
	}
	fmt.Printf("%v", personsArray)
}
