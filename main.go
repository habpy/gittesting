package main

import (
	"fmt"
)

type hub struct {
	achternaam string
	leeftijd   int
	getrouwd   bool
}

func somefunction(a string, b int, g bool) {
	fmt.Printf("The name is %s the age is %d and getrouwd is %t \n", a, b, g)
}

func (h *hub) someOtherFunction() {
	h.achternaam = "Bond"
	h.leeftijd = 40
	h.getrouwd = true

	fmt.Println("The variables are:", h.achternaam, h.leeftijd, h.getrouwd)
}

func main() {
	fmt.Println("This is a Go test")
	h := hub{
		achternaam: "Duran",
		leeftijd:   34,
		getrouwd:   true,
	}
	fmt.Println("The struct is: ", h)
	somefunction("Bond", 40, true)
	h.someOtherFunction()
}
