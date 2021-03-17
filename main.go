package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main function
func main() {

	state := setUp()

	fmt.Println("Welcome to the hamburger shop. Your job is to make a hamburger for a customer. Good luck!")
	var shouldExit bool = false

	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() || !shouldExit {
		s := scanner.Text()

		state = processCommand(s, state)

		fmt.Print(">")
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

}

func processCommand(command string, state State) State {

	verb := findVerb(command, state)
	objectName, objectType := findObject(command, state)

	fmt.Println(verb)
	if verb == "quit" {
		os.Exit(0)
	} else if verb == "get" && objectType == "item" {
		item, _ := findItemByID(objectName, state)
		state.player.items = addItemToInventory(item, state.player.items)
		// state.player.items = append(state.player.items, item)
	} else if verb == "put" && objectType == "item" {
		item, _ := findItemByID(objectName, state)
		state.player.items = RemoveItemFromInventory(item, state.player.items)
	} else if verb == "inventory" {
		fmt.Println(state.player.items)
	} else if verb == "go" && objectType == "room" {
		fmt.Println("add player")
		state.player = addRoomToCharacter(objectName, state)
		// roomToLeave, _ := findRoomByID(state.Player., state)
		roomToGo, _ := findRoomByID(objectName, state.rooms)
		if roomToGo.firstEnter {
			roomToGo.onEnter(roomToGo)
		}

		state.rooms = markRoomAsEntered(objectName, state.rooms)
		// fmt.Println("remove player")
		// state.rooms = removeCharacterFromRoom(state.player, state.rooms)
		// state.rooms = addCharacterToRoom(state.player, roomToGo, state.rooms)
	} else if verb == "where" {
		fmt.Println(state.player.roomID)
	} else if verb == "describe" {
		roomToDescribe, _ := findRoomByID(state.player.roomID, state.rooms)
		fmt.Println(roomToDescribe.description)
	}

	//take order
	//start grill
	//get hamburger patty
	//put hamburger patty on grill
	//get bun
	//put bun on grill
	//get cheese slice
	//put cheese on hamburger patty
	//get plate
	//set down plate
	//put bun on plate
	//take hamburger patty from grill
	//put hamburger patty on bun
	//put onions on hamburger patty
	//put ketchup on hamburger patty
	return state
}

func findObject(command string, state State) (objectName string, objectType string) {

	// fmt.Printf("finding Object from: %v\n", command)
	for _, item := range state.items {
		for _, alias := range item.Aliases {
			if strings.Contains(command, alias) {
				// fmt.Printf("found: %v\n", item.Name)
				return item.Name, "item"
			}
		}
	}

	for _, room := range state.rooms {
		// fmt.Println(verb)
		if strings.Contains(command, room.name) {
			return room.name, "room"
		}
	}

	for _, character := range state.characters {
		// fmt.Println(verb)
		if strings.Contains(command, character.name) {
			return character.name, "character"
		}
	}

	return "", ""

}

func setUp() State {

	var state State

	state.items = createItems()
	state.verbs = createVerbs()
	state.player = createPlayer(state)
	state.rooms = createRooms(state)
	// characters = createCharacters()

	return state
}

// func createCharacters() []Character {

// }

func createNodes(state State) []Node {
	return []Node{
		{
			id:     "action1",
			nextID: "action2",
			successConditions: []Condition{
				{
					completed: false,
					statement: func(state State) bool {
						return state.player.roomID == "living room"
					},
				},
			},
		},
	}
}
