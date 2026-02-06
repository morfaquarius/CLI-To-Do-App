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

func List(tasks []Task, filter string) []Task {
    // Тут будет код
    return tasks // временно
}