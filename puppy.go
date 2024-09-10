package puppy

import (
	"github.com/anuvara2000/dog"
	"fmt"
)

func Bark() string {
	return "Woof!"
}

func Barks() string{
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func Form13(){
	fmt.Println("I'm from version 1.3.0")
}

func Form11(){
	fmt.Println("I'm from version 1.1.0")
}


