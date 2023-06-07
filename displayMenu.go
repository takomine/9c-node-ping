package main

import "fmt"

func displayMenu() {
	fmt.Println("A simple crappy node checker/replacer for Nine Chronicles")
	fmt.Println()
	fmt.Println("Enter the number of your corresponding choice")
	fmt.Println()
	fmt.Println("1 -> Connect Nine Chronicles to a specific node")
	fmt.Println("2 -> Check Ping")
	fmt.Println("3 -> Reset config.json")
	fmt.Println()
	fmt.Println("4 -> Exit")
	fmt.Println()
	fmt.Print(">  ")
}
