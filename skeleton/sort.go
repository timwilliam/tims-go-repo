package main

import (
	"fmt"
	"sort"
)

func main() {
	students := map[string]int{
		"Andi":  10,
		"Citra": 50,
		"Budi":  10,
		"Coco":  40,
	}

	type nameScore struct {
		name  string
		score int
	}

	var studentsSlice []nameScore
	for name, score := range students {
		studentsSlice = append(studentsSlice, nameScore{name, score})
	}

	// first sort by score, then by name (multiple paramaters)
	sort.Slice(studentsSlice, func(i, j int) bool {
		if studentsSlice[i].score > studentsSlice[j].score {
			return true
		} else if studentsSlice[i].score < studentsSlice[j].score {
			return false
		}
		return studentsSlice[i].name < studentsSlice[j].name
	})

	for i := 0; i < len(studentsSlice); i++ {
		fmt.Printf("%s %d\n", studentsSlice[i].name, studentsSlice[i].score)
	}
}
