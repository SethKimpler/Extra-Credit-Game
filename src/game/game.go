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
	win bool
}

type player struct {
	name string
	health int
	tool string
}

type enemy struct {
	name string
	description string
	weapon string
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
	one := room{identifier: 1, name: "One", description: "", up: nil, down: nil, left: nil, right: nil}
	two := room{2, "Two", "", nil, nil, nil, nil}
	three := room{3, "Three", "", nil, nil, &two, nil}
	four := room{4, "Four", "", nil, nil, &three, nil}
	five := room{5, "Five", "", &two, nil, &one, nil}
	six := room{6, "Six", "", &three, nil, &five, nil}
	seven := room{7, "Seven", "", &four, nil, &six, nil}
	eight := room{8, "Eight", "", &five, nil, nil, nil}
	nine := room{9, "Nine", "", &six, nil, &eight, nil}
	ten := room{10, "Ten", "", &seven, nil, &nine, nil}
	c := one
	c.right = &five
	one.right = &two
	one.identifier = 1
	two.identifier = 2
	three.identifier = 3
	four.identifier = 4
	five.identifier = 5
	six.identifier = 6
	seven.identifier = 7
	eight.identifier = 8
	nine.identifier = 9
	ten.identifier = 10
	return game{one, two, three, four, five, six, seven, eight, nine, ten, c, character, false}
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

	character := player{name, 100, ""}
	fmt.Println("Congratulations on creating your character!")
	return character
}

func playGame(g game) {
	scanner := bufio.NewReader(os.Stdin)
	for g.win == false {
		fmt.Println("Current Room: " + g.currentRoom.name)
		fmt.Println("Description: " + g.currentRoom.description)
		fmt.Println("")
		fmt.Println("Exits: ")
		if g.currentRoom.up != nil {
			fmt.Println("Up-" + g.currentRoom.up.name)
		}
		if g.currentRoom.right != nil {
			fmt.Println("Right-" + g.currentRoom.right.name)
		}
		if g.currentRoom.down != nil {
			fmt.Println("Down-" + g.currentRoom.down.name)
		}
		if g.currentRoom.left != nil {
			fmt.Println("Left-" + g.currentRoom.left.name)
		}

		fmt.Print("Enter your next move: ")
		input, _ := scanner.ReadString('\n')
		switch input {
		case "U":
			if g.currentRoom.up != nil {
				//How do I make current room point to this?
				g.currentRoom = *g.currentRoom.up
			}
		case "R":
			if g.currentRoom.right != nil {
				g.currentRoom = *g.currentRoom.right
			}
		case "L":
			if g.currentRoom.left != nil {
				g.currentRoom = *g.currentRoom.left
			}
		case "D":
			if g.currentRoom.down != nil {
				g.currentRoom = *g.currentRoom.down
			}
		}
		fmt.Println("------------")
	}
}

func main() {
	fmt.Println("\n\n")
	character := loadPlayer()
	game := loadGame(character)
	playGame(game)
}
