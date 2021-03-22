package main

import (
	"errors"
	"fmt"
)

func createRooms(state State) []Room {

	var roomsToReturn []Room
	roomsToReturn = append(roomsToReturn, Room{
		name:        "kitchen",
		firstEnter:  true,
		onEnter:     enterKitchen,
		description: "The kitchen is pretty clean. It looks like there is a stove a microwave on a checkered floor."})

	roomsToReturn = append(roomsToReturn, Room{
		name:        "living room",
		firstEnter:  true,
		onEnter:     enterLivingRoom,
		description: "The living room looks dated... The couch is plaid and wrapped in plastic. There is an endtable with a candy dish on it."})

	roomsToReturn = append(roomsToReturn, Room{
		name:        "outside",
		firstEnter:  true,
		onEnter:     enterOutside,
		description: "It is a warm day. You can feel the sun on your face.",
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

func addRoomToCharacter(roomID string, state State) Character {
	state.player.roomID = roomID
	return state.player
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

func findRoomByID(roomIdToFind string, rooms []Room) (Room, error) {

	for _, room := range rooms {
		if room.name == roomIdToFind {
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
	fmt.Printf("first enter %s\n", state.player.roomID)
	//do something else
	return state
}

func enterLivingRoom(state State) State {
	fmt.Printf("first enter %s\n", state.player.roomID)
	//do something else
	return state
}

func enterOutside(state State) State {
	fmt.Printf("first enter %s\n", state.player.roomID)
	//do something else
	return state
}
