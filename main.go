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
		if isItemAvailable(objectName, state) {
			state.player.items = addItemToInventory(Item{name: objectName, amount: 1}, state.player.items)
		} else {
			fmt.Printf("%s is not available to pickup.\n", objectName)
		}
	} else if verb == "put" && objectType == "item" {
		// item, _ := findItemByID(objectName, state)
		state.player.items = removeItemFromInventory(Item{name: objectName, amount: 1}, state.player.items)
	} else if verb == "inventory" {
		fmt.Println(state.player.items)
	} else if verb == "go" && objectType == "room" {
		roomToGo, error := findRoomByID(objectName, state.rooms)
		if error != nil {
			fmt.Printf("There is not room available called %s", objectName)
		} else if state.player.roomID == objectName {
			fmt.Printf("You are already in the %s\n", objectName)
		} else {
			if isRoomAvailable(objectName, state) {
				state = addRoomToCharacter(objectName, state)

				if roomToGo.firstEnter {
					state = roomToGo.onEnter(state)
				}

				state.rooms = markRoomAsEntered(objectName, state.rooms)
			} else {
				fmt.Printf("You are not able to travel to %s\n", objectName)
			}
		}
	} else if verb == "go" {
		fmt.Println("That location is not available.")

	} else if verb == "where" {
		fmt.Println(state.player.roomID)
	} else if verb == "describe" {
		roomToDescribe, _ := findRoomByID(state.player.roomID, state.rooms)
		fmt.Println(roomToDescribe.description)
	} else if verb == "tasks" {
		for _, task := range state.player.taskIDs {
			var fullTask, _ = getTask(task, state)
			fmt.Println(fullTask.successCondition.description)
		}
		// fmt.Println(state.player.taskIDs)
	}

	state = checkTasks(state)

	return state
}

func findObject(command string, state State) (objectName string, objectType string) {

	for _, item := range state.itemsDefinitions {
		for _, alias := range item.aliases {
			if strings.Contains(command, alias) {
				// fmt.Printf("found: %v\n", item.Name)
				return item.name, "item"
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

// func findAmount(command string, state State) int {

// }

func setUp() State {

	var state State

	state.itemsDefinitions = createItems()
	state.verbs = createVerbs()
	state.player = createPlayer(state)
	state.rooms = createRooms(state)
	state.tasks = createTasks(state)

	return state
}

// func createCharacters() []Character {

// }
