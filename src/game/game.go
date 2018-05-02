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
	object string
	foe enemy
}

func loadGame(character player) game {
	one := room{identifier: 1, name: "One", description: "", up: nil, down: nil, left: nil, right: nil, object: "", foe: enemy{}}
	two := room{2, "Two", "", nil, nil, nil, nil, "", enemy{}}
	three := room{3, "Three", "", nil, nil, &two, nil, "", enemy{}}
	four := room{4, "Four", "", nil, nil, &three, nil, "", enemy{}}
	five := room{5, "Five", "", &two, nil, nil, nil, "", enemy{}}
	six := room{6, "Six", "", &three, nil, &five, nil, "", enemy{}}
	seven := room{7, "Seven", "", &four, nil, &six, nil, "", enemy{}}
	eight := room{8, "Eight", "", &five, nil, nil, nil, "", enemy{}}
	nine := room{9, "Nine", "", &six, nil, &eight, nil, "", enemy{}}
	ten := room{10, "Ten", "", &seven, nil, &nine, nil, "", enemy{}}
	c := room{}

	one.identifier = 1
	one.description = "Professor Waldon's Office"
	one.right = &five

	two.identifier = 2
	two.description = "Lobby of Anne Belk Library"
	two.right = &three
	two.down = &five
	two.object = "Sheer Willpower"

	three.identifier = 3
	three.description = "1st floor of Anne Belk"
	three.right = &four
	three.down = &six
	three.foe = enemy{"Joel Swanson", "Discrete Math Professor", "i-clicker quizzes"}

	four.identifier = 4
	four.description = "Bell outside of Anne Belk"
	four.down = &seven
	four.object = "Exception Handler"

	five.identifier = 5
	five.description = "Long, dark hallway in the second floor of Anne Belk"
	five.right = &six
	five.down = &eight

	six.identifier = 6
	six.description = "Sandford Mall, a vast, beatiful oasis"
	six.right = &seven
	six.down = &nine

	seven.identifier = 7
	seven.description = "Sanford building, home of all RC 1000 classes"
	seven.down = &ten

	eight.identifier = 8
	eight.description = "Kidd Brewer Stadium"
	eight.right = &nine
	eight.foe = enemy{"Java Compiler", "Enemy of all things good", "Null-Pointer Exception"}

	nine.identifier = 9
	nine.description = "Varsity gym. #gainz"
	nine.right = &ten
	nine.object = "Quiz Master Badge"

	ten.identifier = 10
	ten.description = "Central Dining Hall, where you say you'll eat healthy, but we all know you're gonna get Chick-fil-a"
	ten.foe = enemy{"Checkstyle", "Grammar Nazi", " 'file must end with new-line' "}

	c = one
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
	fmt.Println("")

	character := player{name, 100, ""}
	fmt.Println("Congratulations on creating your character!")
	fmt.Println("")
	fmt.Println("------------")
	return character
}

func printMap() {
	fmt.Println("")
	fmt.Println("    2 - 3 - 4")
	fmt.Println("    |   |   |")
	fmt.Println("1 - 5 - 6 - 7")
	fmt.Println("    |   |   |")
	fmt.Println("    8 - 9 -10")
	fmt.Println("")
}

func playGame(g game) {
	scanner := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Current Room: " + g.currentRoom.name)
		fmt.Println("Description: " + g.currentRoom.description)
		fmt.Println("Tools in room: " + g.currentRoom.object)
		fmt.Println("")
		fmt.Println("Exits: ")
		if g.currentRoom.up != nil {
			fmt.Println("Direction: Up / Name: " + g.currentRoom.up.name)
		}
		if g.currentRoom.right != nil {
			fmt.Println("Direction: Right / Name: " + g.currentRoom.right.name)
		}
		if g.currentRoom.down != nil {
			fmt.Println("Direction: Down / Name: " + g.currentRoom.down.name)
		}
		if g.currentRoom.left != nil {
			fmt.Println("Direction: Left / Name: " + g.currentRoom.left.name)
		}
		fmt.Println("")

		if g.currentRoom.foe.name != "" {
			fmt.Println("ENEMY IN ROOM!")
			fmt.Println("Name: " + g.currentRoom.foe.name)
			fmt.Println("Description: " + g.currentRoom.foe.description)
			fmt.Println("Weapon: " + g.currentRoom.foe.weapon)
			fmt.Println("")
		}

		fmt.Println("Use R, L, U, D to move")
		fmt.Println("Use F to fight (if enemy is present)")
		fmt.Println("Use P to pick-up object, M to display map \n and C to display your character's attributes")

		fmt.Print("Enter your next move: ")
		input, _ := scanner.ReadString('\n')
		switch input {
		case "U\n":
			if g.currentRoom.up != nil {
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
		case "M\n":
			printMap()
		case "P\n":
			if g.currentRoom.object != "" {
				g.currentPlayer.tool = g.currentRoom.object
				g.currentRoom.object = ""
			}
		case "C\n":
			fmt.Println("")
			fmt.Print("Name: " + g.currentPlayer.name)
			fmt.Print("Health: ")
			fmt.Println(g.currentPlayer.health)
		case "F\n":
			if g.currentPlayer.tool == "" {
				fmt.Println("You have no tool, are you sure you would like to fight?")
			} else {
				fmt.Println("Are you sure you want to fight " + g.currentRoom.foe.name + " with " + g.currentPlayer.tool + "?")
			}
			fmt.Println("Enter Y or N")
			answer, _ := scanner.ReadString('\n')
			if answer == "Y\n" {
				switch g.currentRoom.foe.name {
				case "Joel Swanson":
					if g.currentPlayer.tool == "Quiz Master Badge" {
						fmt.Println("VICTORY!")
						g.three.foe.name = ""
					} else {
						fmt.Println("DEFEAT!")
						g.currentPlayer.health -= 34
					}
				case "Java Compiler":
					if g.currentPlayer.tool == "Exception Handler" {
						fmt.Println("VICTORY!")
						g.eight.foe.name = ""
					} else {
						fmt.Println("DEFEAT!")
						g.currentPlayer.health -= 34
					}
				case "Checkstyle":
					if g.currentPlayer.tool == "Sheer Willpower" {
						fmt.Println("VICTORY (somehow)!")
						g.ten.foe.name = ""
					} else {
						fmt.Println("DEFEAT!")
						g.currentPlayer.health -= 34
					}
				}
			}
		}

		if g.three.foe.name == "" && g.eight.foe.name == "" && g.ten.foe.name == "" {
			fmt.Println("CONGRATULATIONS!!! You have beaten 'Waldon's World'!!!")
			break
		}

		if g.currentPlayer.health < 0 {
			fmt.Println("Your player has died... \nDEFEAT")
			break
		}
		fmt.Println("")
		fmt.Println("------------")
	}
}

func main() {
	fmt.Println("\n")
	character := loadPlayer()
	fmt.Print("The Appalachian State campus needs your help! ")
	fmt.Print("The Java-bot and Rogue Professors have taken over, and it is your job to save us! ")
	fmt.Print("By searching around campus, you will be able to find various tools that you can use to help free the campus! ")
	fmt.Print("But be wary, these enemies will be hiding where you lease expect it. ")
	fmt.Print("Use your tools against the enemy they are best suited for, and you can defeat them. ")
	fmt.Println("Defeat each enemy, and you will be remembered as a hero! Good Luck, and may the compiler be ever in your favor! ")
	game := loadGame(character)
	playGame(game)
}

