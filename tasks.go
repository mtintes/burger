package main

import (
	"errors"
	"fmt"
)

func nextTask(currentTask Task, tasks []Task) Task {
	for _, action := range tasks {
		if action.id == currentTask.nextID {
			return action
		}
	}
	return currentTask
}

func getTask(taskID string, state State) (Task, error) {
	for _, task := range state.tasks {
		if task.id == taskID {
			return task, nil
		}
	}

	return Task{id: "none"}, errors.New("NOT_FOUND")

}

func createTasks(state State) []Task {
	return []Task{
		{
			id:     "action1",
			nextID: "action2",
			successConditions: []Condition{
				{
					completed: false,
					statement: func(state State) bool {
						return state.player.roomID == "living room"
					},
				},
			},
		},
		{
			id:     "action2",
			nextID: "action3",
			successConditions: []Condition{
				{
					completed: false,
					statement: func(state State) bool {
						fmt.Println("You walked into the kitchen")
						return state.player.roomID == "kitchen"
					},
				},
			},
		},
		{
			id:     "action3",
			nextID: "complete",
			successConditions: []Condition{
				{
					completed: false,
					statement: func(state State) bool {
						fmt.Println("You need another item")
						return len(state.player.items) > 2
					},
				},
			},
		},
	}
}

func checkTasks(state State) State {
	//checks to see if tasks are completed

	var updatedTasks []string

	if len(state.player.taskIDs) > 0 {
		for _, taskName := range state.player.taskIDs {
			updatedTasks = append(updatedTasks, checkTask(taskName, state))
		}
	}

	state.player.taskIDs = updatedTasks

	return state

}

func checkTask(taskID string, state State) string {

	fmt.Println("TaskToCheck:", taskID)
	var taskToCheck, _ = getTask(taskID, state)

	for _, condition := range taskToCheck.successConditions {
		if condition.statement(state) {
			fmt.Println("You did a thing")
			return checkTask(taskToCheck.nextID, state)
			// update the task list to the new task
		} else {
			return taskID
		}
	}

	return ""
}
