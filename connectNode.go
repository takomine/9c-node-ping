package main

import (
	"fmt"
	"strconv"
)

func connectNode() {
	for {
		fmt.Println("What node do you want to use?")
		fmt.Println()
		fmt.Println("1 -> 9c-main-rpc-1.nine-chronicles.com (RPC Node 1)")
		fmt.Println("2 -> 9c-main-rpc-2.nine-chronicles.com (RPC Node 2)")
		fmt.Println("3 -> 9c-main-rpc-3.nine-chronicles.com (RPC Node 3)")
		fmt.Println("4 -> 9c-main-rpc-4.nine-chronicles.com (RPC Node 4)")
		fmt.Println("5 -> 9c-main-rpc-5.nine-chronicles.com (RPC Node 5)")
		fmt.Println("6 -> tky-nc-1.ninodes.com (Tokyo Node)")
		fmt.Println("7 -> sgp-nc-1.ninodes.com (Singapore Node)")
		fmt.Println("8 -> nj-nc-1.ninodes.com (New Jersey Node)")
		fmt.Println("9 -> la-nc-1.ninodes.com (Los Angeles Node)")
		fmt.Println("10 -> fra-nc-1.ninodes.com (France Node)")
		fmt.Println()
		fmt.Println("11 -> Go back")
		fmt.Println()
		fmt.Print("> ")

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
			node := "9c-main-rpc-1.nine-chronicles.com,80,31238"
			overwriteJson(node)
			return
		case 2:
			clearScreen()
			node := "9c-main-rpc-2.nine-chronicles.com,80,31238"
			overwriteJson(node)
			return
		case 3:
			clearScreen()
			node := "9c-main-rpc-3.nine-chronicles.com,80,31238"
			overwriteJson(node)
			return
		case 4:
			clearScreen()
			node := "9c-main-rpc-4.nine-chronicles.com,80,31238"
			overwriteJson(node)
			return
		case 5:
			clearScreen()
			node := "9c-main-rpc-5.nine-chronicles.com,80,31238"
			overwriteJson(node)
			return
		case 6:
			clearScreen()
			node := "tky-nc-1.ninodes.com,80,31238"
			overwriteJson(node)
			return
		case 7:
			clearScreen()
			node := "sgp-nc-1.ninodes.com,80,31238"
			overwriteJson(node)
			return
		case 8:
			clearScreen()
			node := "nj-nc-1.ninodes.com,80,31238v"
			overwriteJson(node)
			return
		case 9:
			clearScreen()
			node := "la-nc-1.ninodes.com,80,31238"
			overwriteJson(node)
			return
		case 10:
			clearScreen()
			node := "fra-nc-1.ninodes.com,80,31238"
			overwriteJson(node)
			return
		case 11:
			clearScreen()
			fmt.Println("Exiting...")
			return
		default:
			clearScreen()
			fmt.Println("Invalid choice. Please try again.")

		}
	}
}
