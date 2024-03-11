package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

// create the function changeTheIdentity here...
func changeTheIdentity(person1 *Person, person2 *Person) {
	tempPerson := *person1;
	*person1 = *person2;
	*person2 = tempPerson;
	return;
}
func main() {
	person1 := Person{firstName: "John", lastName: "Doe", age: 20}
	person2 := Person{firstName: "James", lastName: "carry", age: 20}
	
	fmt.Printf("Before Swapping  \n person1 %v \n person2 %v \n",person1,person2);
	changeTheIdentity(&person1,&person2);
	fmt.Printf("After Swapping  \n person1 %v \n person2 %v \n",person1,person2);
}
