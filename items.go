package main

import (
	"errors"
	"fmt"
)

func createItems() []ItemDefinition {
	return []ItemDefinition{
		{
			name:    "hamburger patty",
			aliases: []string{"burger", "hamburger", "patty", "hamburger patty", "burgers", "patties", "hamburgers"},
		},
		{
			name:    "bun",
			aliases: []string{"bun", "bread"},
		},
		{
			name:    "hat",
			aliases: []string{"hat"},
		},
	}
}

func addItemToInventory(itemToAdd Item, items []Item) []Item {
	var alreadyHasItem = false

	var updatedItems []Item
	for _, item := range items {
		if item.name == itemToAdd.name {
			alreadyHasItem = true
			updatedItem := Item{name: item.name, amount: item.amount + itemToAdd.amount}
			updatedItems = append(updatedItems, updatedItem)
		} else {
			updatedItems = append(updatedItems, item)
		}
	}

	if alreadyHasItem == false {
		itemToAdd.amount = 1
		updatedItems = append(updatedItems, itemToAdd)
	}

	return updatedItems
}

func isItemAvailable(itemID string, state State) bool {
	fmt.Println(itemID)
	room, _ := findRoomByID(state.player.roomID, state.rooms)
	fmt.Println(room)
	if contains(itemID, room.items) {
		return true
	}
	return false
}

func contains(itemID string, items []Item) bool {
	for _, item := range items {
		if itemID == item.name {
			return true
		}
	}
	return false
}

func removeItemFromInventory(itemToRemove Item, items []Item) []Item {
	var updatedItems []Item
	// fmt.Printf("looking for: %v\n", itemToRemove.Name)
	for _, item := range items {
		// fmt.Printf("looked at %v\n", item.Name)
		if item.name != itemToRemove.name {
			updatedItems = append(updatedItems, item)
		} else if item.name == itemToRemove.name && item.amount > itemToRemove.amount {
			item.amount = item.amount - itemToRemove.amount
			updatedItems = append(updatedItems, item)
		}
	}

	return updatedItems
}

func findItemByID(id string, state State) (ItemDefinition, error) {
	for _, item := range state.itemsDefinitions {
		for _, alias := range item.aliases {
			if id == alias {
				return item, nil
			}
		}
	}

	return ItemDefinition{name: "none"}, errors.New("No Item Found")
}
