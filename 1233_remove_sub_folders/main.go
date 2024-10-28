package main

import (
	"fmt"
	"sort"
	"strings"
)

// removeSubfolders returns all folders that are removed sub folders inside
// Complexity:
// m: length of folder array
// n: length of longest number of folders name
// O(m x n)
// Space: O(m)
func removeSubfolders(folders []string) []string {
	visited := make(map[string]bool)
	var result []string

	sort.Slice(folders, func(i, j int) bool {
		return len(folders[i]) < len(folders[j])
	})

	for _, folder := range folders {
		if isParent(visited, folder) {
			result = append(result, folder)
		}
	}

	return result
}

func isParent(visited map[string]bool, folder string) bool {
	parts := strings.Split(folder, "/")
	if len(parts) == 0 {
		return false
	}

	cFolder := ""
	cFolder = parts[0]
	for i := 1; i < len(parts); i++ {
		cFolder = fmt.Sprintf("%s/%s", cFolder, parts[i])
		if _, exists := visited[cFolder]; exists {
			return false
		}

	}

	visited[folder] = true
	return true
}

func main() {
	folders := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	fmt.Println(removeSubfolders(folders))
}
