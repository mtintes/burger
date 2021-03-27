package main

func createPlayer(state State) Character {
	var player = Character{
		name: "You",
	}

	// hat, _ := findItemByID("hat", state)
	player.items = addItemToInventory(Item{name: "hat", amount: 1}, player.items)
	player.roomID = "kitchen"
	player.taskIDs = []string{"t1a1", "t2a1"}
	return player
}
