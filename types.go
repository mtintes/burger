package main

// Item is
type Item struct {
	Name    string
	Aliases []string
	Amount  int
}

// Place is a space in a room that a character can go. ex. cabinets(space) in a kitchen(room) or drawers(space) in a bedroom(room)
type Place struct {
	name       string
	characters []Character
	items      []Item
}

// Room is a space that a character can go. ex. kitchen or hallway
type Room struct {
	name        string
	characters  []Character
	onEnter     func(state State) State
	description string
	firstEnter  bool
}

//Character is
type Character struct {
	name    string
	items   []Item
	roomID  string
	taskIDs []string
}

//Verb is
type Verb struct {
	primary string
	aliases []string
}

//State is
type State struct {
	rooms      []Room
	items      []Item
	characters []Character
	verbs      []Verb
	player     Character
	tasks      []Task
}

//Task is
type Task struct {
	id                string
	nextID            string
	successConditions []Condition
}

//Condition is
type Condition struct {
	completed bool
	statement func(state State) bool
}

// // Room is
// type Room struct {
// 	name       string
// 	characters []Character
// 	onEnter    func()
// }

// // Item is
// type Item struct {
// 	name    string
// 	aliases []string
// }

// //Character is
// type Character struct {
// 	name  string
// 	items []Item
// }

// //Verb is
// type Verb struct {
// 	primary string
// 	aliases []string
// }

// //State is
// type State struct {
// 	rooms     []Room
// 	items     []Item
// 	character []Character
// 	verbs     []Verb
// 	player    Character
// }
