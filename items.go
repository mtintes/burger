package main

import "errors"

func createItems() []Item {
	return []Item{
		{
			Name:    "hamburger patty",
			Aliases: []string{"burger", "hamburger", "patty", "hamburger patty", "burgers", "patties", "hamburgers"},
		},
		{
			Name:    "bun",
			Aliases: []string{"bun", "bread"},
		},
		{
			Name:    "hat",
			Aliases: []string{"hat"},
		},
	}
}

func addItemToInventory(itemToAdd Item, items []Item) []Item {
	var alreadyHasItem = false

	var newItems []Item
	// fmt.Printf("looking for: %v\n", itemToAdd.Name)
	for _, item := range items {
		// fmt.Printf("looked at %v\n", item.Name)
		if item.Name == itemToAdd.Name {
			alreadyHasItem = true
			newItem := Item{Name: item.Name, Aliases: item.Aliases, Amount: item.Amount + 1}
			newItems = append(newItems, newItem)
		} else {
			newItems = append(newItems, item)
		}
	}

	if alreadyHasItem == false {
		itemToAdd.Amount = 1
		newItems = append(newItems, itemToAdd)
	}

	return newItems
}

func RemoveItemFromInventory(itemToRemove Item, items []Item) []Item {
	var newItems []Item
	// fmt.Printf("looking for: %v\n", itemToRemove.Name)
	for _, item := range items {
		// fmt.Printf("looked at %v\n", item.Name)
		if item.Name != itemToRemove.Name {
			newItems = append(newItems, item)
		} else if item.Name == itemToRemove.Name && item.Amount > 1 {
			item.Amount = item.Amount - 1
			newItems = append(newItems, item)
		}
	}

	return newItems

}

func findItemByID(id string, state State) (Item, error) {
	for _, item := range state.items {
		for _, alias := range item.Aliases {
			if id == alias {
				return item, nil
			}
		}
	}

	return Item{Name: "none"}, errors.New("No Item Found")
}
