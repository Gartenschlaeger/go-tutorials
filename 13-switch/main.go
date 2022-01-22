package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Value of dice is %v\n", diceNumber)

	// In Go wird automatisch immer nur der aktuelle case bearbeitet,
	// break ist somit nicht erforderlich!

	// Um von einem case zum nächsten weiterzugehen,
	// kann das Schlüsselwort fallthrough verwendet werden!

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move to 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")
	}

}
