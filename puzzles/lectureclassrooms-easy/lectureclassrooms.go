package main

import (
	"fmt"
	"sort"
)

/**
Given an array of time intervals (start, end) for classroom lectures
(possibly overlapping), find the minimum number of rooms required.

For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.
*/

type class struct {
	start, end int
}

func main() {
	classes := setupClasses()

	// sort the classes by start time
	sort.Slice(classes, func(i, j int) bool { //[class1, class0, class2]
		return classes[i].start < classes[j].start
	})

	//go through the sorted classes and add rooms
	rooms := make([]class, 0)
	for _, cl := range classes {
		if len(rooms) == 0 {
			rooms = append(rooms, cl)
			continue
		}
		var reusedRoom bool
		// there are some rooms in progress already let's try to reuse them
		for o, ongoing := range rooms {
			// the ongoing class has to end before the new class starts
			if ongoing.end < cl.start {
				// we put in the new class in the old room
				rooms[o] = cl
				reusedRoom = true
				break
			}
		}
		if !reusedRoom {
			rooms = append(rooms, cl)
		}
		// sort the rooms by end time again as we want to
		// get the first ending class first to start classes as soon as possible
		sort.SliceStable(rooms, func(i, j int) bool {
			return rooms[i].end < rooms[j].end
		})
	}

	fmt.Printf("Minimum rooms: %d \n", len(rooms))
}
func setupClasses() []class {
	return []class{
		{
			start: 15,
			end:   45,
		},
		{
			start: 30,
			end:   75,
		},
		{
			start: 0,
			end:   50,
		},
		{
			start: 60,
			end:   150,
		},
		{
			start: 60,
			end:   120,
		},
	}
}
