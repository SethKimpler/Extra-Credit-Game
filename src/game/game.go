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
	five := room{5, "Five", "", &two, nil, nil, nil}
	six := room{6, "Six", "", &three, nil, &five, nil}
	seven := room{7, "Seven", "", &four, nil, &six, nil}
	eight := room{8, "Eight", "", &five, nil, nil, nil}
	nine := room{9, "Nine", "", &six, nil, &eight, nil}
	ten := room{10, "Ten", "", &seven, nil, &nine, nil}
	c := room{}

	one.identifier = 1
	one.description = "Room one"
	one.right = &five

	two.identifier = 2
	two.description = "Room two"
	two.right = &three
	two.down = &five

	three.identifier = 3
	three.description = "Room three"
	three.right = &four
	three.down = &six

	four.identifier = 4
	four.description = "Room four"
	four.down = &seven

	five.identifier = 5
	five.description = "Room five"
	five.right = &six
	five.down = &eight

	six.identifier = 6
	six.description = "Room six"
	six.right = &seven
	six.down = &nine

	seven.identifier = 7
	seven.description = "Room seven"
	seven.down = &ten

	eight.identifier = 8
	eight.description = "Room eight"
	eight.right = &nine

	nine.identifier = 9
	nine.description = "Room nine"
	nine.right = &ten

	ten.identifier = 10
	ten.description = "Room ten"
	c = one
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
	fmt.Println("")

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
		fmt.Println("")
		fmt.Println("Use R, L, U, D to move")

		fmt.Print("Enter your next move: ")
		input, _ := scanner.ReadString('\n')
		switch input {
		case "U\n":
			if g.currentRoom.up != nil {
				//How do I make current room point to this?
				g.currentRoom = *g.currentRoom.up
			}
		case "R\n":
			if g.currentRoom.right != nil {
				g.currentRoom = *g.currentRoom.right
			}
		case "L\n":
			if g.currentRoom.left != nil {
				g.currentRoom = *g.currentRoom.left
			}
		case "D\n":
			if g.currentRoom.down != nil {
				g.currentRoom = *g.currentRoom.down
			}
		case "?\n":
			fmt.Println("R = move right, L = move left, U = move up, D = move down")
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
