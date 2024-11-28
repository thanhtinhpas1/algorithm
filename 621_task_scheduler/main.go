package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	frequencyMap := make(map[byte]int)
	var performed []byte

	for _, task := range tasks {
		frequencyMap[task]++
	}

	for len(frequencyMap) > 0 {
		maxFre := -1
		selectTask := byte('.')

		for task, frequency := range frequencyMap {
			if !didPerformed(performed, task, n) {
				if frequency > maxFre {
					maxFre = frequency
					selectTask = task
				}
			}
		}

		if selectTask == byte('.') {
			performed = append(performed, byte('.'))
			continue
		}

		performed = append(performed, selectTask)
		frequencyMap[selectTask]--

		if frequencyMap[selectTask] == 0 {
			delete(frequencyMap, selectTask)
		}

	}

	return len(performed)
}

func didPerformed(performed []byte, task byte, n int) bool {
	i := 0
	if len(performed) > n {
		i = len(performed) - n
	}

	for i < len(performed) {
		if task == performed[i] {
			return true
		}
		i++
	}

	return false
}

func main() {
	tasks := []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'}
	fmt.Println(leastInterval(tasks, 1))
}
