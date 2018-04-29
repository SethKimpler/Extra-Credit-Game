package main

import "fmt"
import "bufio"
import "os"

type game struct {
	one room
	two room
	three room
	four room
	five room
	six room
	seven room
	eight room
	nine room
	ten room
	currentRoom room
	currentPlayer player
}

type player struct {
	name string
	health int
}

type room struct {
	identifier int
	name string
	description string
	up *room
	down *room
	left *room
	right *room
}

func loadGame(character player) game {
	one := room{}
	one.identifier = 1
	two := room{}
	two.identifier = 2
	three := room{}
	three.identifier = 3
	four := room{}
	four.identifier = 4
	five := room{}
	five.identifier = 5
	six := room{}
	six.identifier = 6
	seven := room{}
	seven.identifier = 7
	eight := room{}
	eight.identifier = 8
	nine := room{}
	nine.identifier = 9
	ten := room{}
	ten.identifier = 10
	c := one
	return game{one, two, three, four, five, six, seven, eight, nine, ten, c, character}
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
	game := loadGame(character)
	fmt.Println(game.currentRoom)
}
