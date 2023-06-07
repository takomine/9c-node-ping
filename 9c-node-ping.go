package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		clearScreen()
		displayMenu()

		//show error if the user entered an invalid choice
		var input string
		fmt.Scanln(&input)

		choices, err := strconv.Atoi(input)
		if err != nil {
			clearScreen()
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		//menu choices
		switch choices {
		case 1:
			clearScreen()
			connectNode()
		case 2:
			clearScreen()
			checkPing()
		case 3:
			clearScreen()
			resetConfig()
		case 4:
			clearScreen()
			fmt.Println("Exiting...")
			return
		/*case 5:
		//test shit
		clearScreen()
		test()*/
		default:
			clearScreen()
			fmt.Println("Invalid choice. Please try again.")

		}
	}
}
