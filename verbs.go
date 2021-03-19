package main

import "strings"

func createVerbs() []Verb {
	return []Verb{
		{
			primary: "get",
			aliases: []string{"get", "pickup", "grab", "take"},
		},
		{
			primary: "put",
			aliases: []string{"put", "set", "drop"},
		},
		{
			primary: "combine",
			aliases: []string{"combine", "mix"},
		},
		{
			primary: "go",
			aliases: []string{"go", "move"},
		},
		{
			primary: "quit",
			aliases: []string{"quit", "exit"},
		},
		{
			primary: "inventory",
			aliases: []string{"inventory", "items"},
		},
		{
			primary: "where",
			aliases: []string{"where"},
		},
		{
			primary: "describe",
			aliases: []string{"describe", "description"},
		},
	}
}

func findVerb(command string, state State) string {

	for _, verb := range state.verbs {
		// fmt.Println(verb)
		for _, alias := range verb.aliases {
			if strings.Contains(command, alias) {
				return verb.primary
			}
		}
	}

	return ""

}
