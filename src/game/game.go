package main

import "fmt"
import "bufio"
import "os"

type game struct {
	currentRoom room
	currentPlayer player
}

type player struct {
	name string
	health int
}

type room struct {
	up *room
	down *room
	left *room
	right *room
	description string
	name string
}

func loadPlayer() player {
	fmt.Println("Welcome to Waldon's World!")
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("The first step is to select your character's name. What shall it be?")
	name, _ := scanner.ReadString('\n')
	for {
		fmt.Print("You have chosen: " + name)
		fmt.Println("Type 'yes' to continue or choose another name: ")
		confirm, _ := scanner.ReadString('\n')
		if confirm == "yes\n" {
			break
		} else {
			name = confirm
		}
	}

	character := player{name, 100}
	fmt.Println("Congratulations on creating your character!")
	return character
}

func main() {
	character := loadPlayer()
	fmt.Print(character.name)
}
