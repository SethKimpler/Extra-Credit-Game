package main

import "fmt"

func testCharacter() {
	test := player{}
	test.name = "Waldon"
	fmt.Println(test)
}

func testRoom() {
	test := room{}
	test.name = "name"
	test.description = "description"
	fmt.Println(test)
}
