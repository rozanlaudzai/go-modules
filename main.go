package go_modules

import (
	"fmt"

	"github.com/rozanlaudzai/go-modules/helper"
	"github.com/rozanlaudzai/go-modules/helper/oof"
)

func awalnyaMain() {

	// first commit
	// second commit
	// third commit

	// this is login feature
	// second commit of login feature

	// biar ada tag
	// biar ada tag part 2

	arrPersons := []*helper.Person{
		&helper.Person{
			Name: "Naruto",
			Age:  18,
		},
	}

	for range 100000 {
		for _, person := range arrPersons {
			person.Name = "kintil"
			person.Age = 31
			person.Children[0] = 100
		}
	}

	fmt.Println(arrPersons[0].Children[0])
	oof.PrintArr([]int{1, 2, 3, 4})

}
