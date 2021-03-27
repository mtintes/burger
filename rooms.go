package main

import (
	"errors"
	"fmt"
)

func createRooms(state State) []Room {

	var roomsToReturn []Room
	roomsToReturn = append(roomsToReturn, Room{
		name:           "kitchen",
		firstEnter:     true,
		onEnter:        enterKitchen,
		description:    "The kitchen is pretty clean. It looks like there is a stove a microwave on a checkered floor.",
		available:      true,
		connectedRooms: []string{"living room"},
		items: []Item{
			{name: "hamburger patty", amount: 1}},
	})

	roomsToReturn = append(roomsToReturn, Room{
		name:           "living room",
		firstEnter:     true,
		onEnter:        enterLivingRoom,
		description:    "The living room looks dated... The couch is plaid and wrapped in plastic. There is an endtable with a candy dish on it.",
		available:      false,
		connectedRooms: []string{"kitchen", "outside"},
		items:          []Item{},
	})

	roomsToReturn = append(roomsToReturn, Room{
		name:           "outside",
		firstEnter:     true,
		onEnter:        enterOutside,
		description:    "It is a warm day. You can feel the sun on your face.",
		available:      false,
		connectedRooms: []string{"living room"},
		items:          []Item{},
	})
	// roomToAdd, _ := findRoomByID("kitchen", roomsToReturn)
	// roomsToReturn = addCharacterToRoom(state.player, roomToAdd, roomsToReturn)
	return roomsToReturn
}

// func addCharacterToRoom(characterToAdd Character, roomToAddCharacter Room, rooms []Room) []Room {

// 	var newRooms []Room

// 	for _, room := range rooms {
// 		if room.name == roomToAddCharacter.name {
// 			room.characters = append(room.characters, characterToAdd)
// 			room.onEnter(state)
// 			newRooms = append(newRooms, room)
// 		} else {
// 			newRooms = append(newRooms, room)
// 		}
// 	}

// 	return newRooms

// }

func addRoomToCharacter(roomID string, state State) State {

	var updatedRooms []Room

	for _, room := range state.rooms {
		if room.name == roomID {
			room.available = true
		}
		updatedRooms = append(updatedRooms, room)
	}

	state.rooms = updatedRooms
	state.player.roomID = roomID
	return state
}

func isRoomAvailable(roomID string, state State) bool {

	playerRoom, _ := findRoomByID(state.player.roomID, state.rooms)

	for _, connectedRoom := range playerRoom.connectedRooms {
		if roomID == connectedRoom {
			return true
		}
	}

	room, error := findRoomByID(roomID, state.rooms)
	// fmt.Printf("room: %s\navailable:%v\n", room.name, room.available)
	if error == nil {
		return room.available
	}

	return false
}

//do we want this to look up the room that the chracter is in OR do we make that a seperate action.
func removeCharacterFromRoom(characterToRemove Character, rooms []Room) []Room {

	var newRooms []Room

	for _, room := range rooms {
		newRoom := room
		newRoom.characters = nil
		for _, character := range room.characters {
			fmt.Printf("looked at %v, found at %v", character.name, characterToRemove.name)
			if character.name != characterToRemove.name {
				newRoom.characters = append(room.characters, character)
			}
		}
		newRooms = append(newRooms, newRoom)
	}

	return newRooms

}

func findRoomByID(roomIDToFind string, rooms []Room) (Room, error) {

	for _, room := range rooms {
		if room.name == roomIDToFind {
			return room, nil
		}
	}

	return Room{name: "none"}, errors.New("NOT_FOUND")
}

func markRoomAsEntered(roomID string, rooms []Room) []Room {

	var updateRooms []Room

	for _, room := range rooms {
		if room.name == roomID {
			room.firstEnter = false
		}
		updateRooms = append(updateRooms, room)
	}

	return updateRooms
}

//On Enters Functions
func enterKitchen(state State) State {
	fmt.Printf("You entered the %s for the first time.\n", state.player.roomID)
	//do something else
	return state
}

func enterLivingRoom(state State) State {
	fmt.Printf("You entered the %s for the first time.\n", state.player.roomID)
	//do something else
	return state
}

func enterOutside(state State) State {
	fmt.Printf("You went %s for the first time.\n", state.player.roomID)
	//do something else
	return state
}
