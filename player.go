package main

func createPlayer(state State) Character {
	var player = Character{
		name: "You",
	}

	hat, _ := findItemByID("hat", state)
	player.items = addItemToInventory(hat, player.items)
	player.roomID = "kitchen"
	return player
}
