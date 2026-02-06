package todo

// Add возвращает обновлённый срез с новой задачей
func Add(tasks []Task, desc string) []Task {
	var maxID int
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	newTask := Task{
		ID:          maxID + 1,
		Description: desc,
		Done:        false,
	}
	tasks = append(tasks, newTask)
	return tasks
}

//List возвращает отфильтрованный срез по заданному параметру
func List(tasks []Task, filter string) []Task {
	switch filter {
	case "all":
		return tasks
	case "done":
		doneTasks := []Task{}
		for _, task := range tasks {
			if task.Done == true {
				doneTasks = append(doneTasks, task)
			}
		}
		return doneTasks
	case "pending":
		pendingTasks := []Task{}
		for _, task := range tasks {
			if task.Done == false {
				pendingTasks = append(pendingTasks, task)
			}
		}
		return pendingTasks
	default:
		return []Task{}
	}
}
