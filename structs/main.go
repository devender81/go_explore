package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create a new Person struct instance
	person := Person{
		Name:    "dev",
		Age:     40,
		Country: "India",
	}

	// Access and modify struct fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Country:", person.Country)

	person.Age = 31
	fmt.Println("Updated Age:", person.Age)

	// Create an anonymous struct instance
	anonymous := struct {
		City    string
		Country string
	}{
		City:    "Hyderabad",
		Country: "India",
	}

	fmt.Println("City:", anonymous.City)
	fmt.Println("Country:", anonymous.Country)
}
