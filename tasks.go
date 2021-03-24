package main

import (
	"errors"
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
			id:     "t1a1",
			nextID: "t1a2",
			successCondition: Condition{
				completed:   false,
				description: "Enter the living room.",
				statement: func(state State) bool {
					return state.player.roomID == "living room"
				},
			},
		},
		{
			id:     "t1a2",
			nextID: "complete",
			successCondition: Condition{
				completed:   false,
				description: "Enter the kitchen",
				statement: func(state State) bool {
					return state.player.roomID == "kitchen"
				},
			},
		},
		{
			id:     "t2a1",
			nextID: "t2a2",
			successCondition: Condition{
				completed:   false,
				description: "Pick up more than 1 item",
				statement: func(state State) bool {
					return len(state.player.items) > 1
				},
			},
		},
		{
			id:     "t2a2",
			nextID: "complete",
			successCondition: Condition{
				completed:   false,
				description: "Pick up more than 2 items",
				statement: func(state State) bool {
					return len(state.player.items) > 2
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

	// fmt.Println("TaskToCheck:", taskID)
	var taskToCheck, _ = getTask(taskID, state)
	// fmt.Println(taskToCheck)

	if taskToCheck.successCondition.statement(state) {
		// fmt.Println("You did a thing")
		if taskToCheck.nextID != "complete" {
			return checkTask(taskToCheck.nextID, state)
		}
		// update the task list to the new task
	}
	return taskID

}
