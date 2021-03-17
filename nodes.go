package main

func nextNode(currentNode Node, nodes []Node) Node {
	for _, action := range nodes {
		if action.id == currentNode.nextID {
			return action
		}
	}
	return currentNode
}
