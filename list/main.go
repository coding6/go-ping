package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskId   int
	TaskBody string
}

var tasks []Task
var lastIdx int

func add(task *Task) {
	if task == nil {
		return
	}
	tasks = append(tasks, *task)
	fmt.Println("task is added in List", tasks)
}

func del(taskId int) {
	for i, task := range tasks {
		if task.TaskId == taskId {
			tasks = append(tasks[i:], tasks[i+1:]...)
			fmt.Println("taskId:", taskId, "is deleted")
			return
		}
	}
	fmt.Println("not found taskId:", taskId)
}

func list() {
	fmt.Println(tasks)
}

func main() {
	lastIdx = 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		switch num {
		case 1:
			fmt.Println("Enter taskBody:")
			scanner.Scan()
			text := scanner.Text()
			lastIdx++
			task := Task{TaskId: lastIdx, TaskBody: text}
			add(&task)
		case 2:
			fmt.Println("Enter taskId:")
			scanner.Scan()
			taskId := scanner.Text()
			intTaskId, _ := strconv.Atoi(taskId)
			del(intTaskId)
		case 3:
			list()
		case 4:
			fmt.Println("exit...")
		}
		fmt.Println()
	}

}
