package go_modules

import "fmt"

type Person struct {
	Name string
	Age  int
}

func SayHello(person Person) {
	fmt.Printf("Hello, %v\n", person.Name)
}
