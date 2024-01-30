package twobestnonoverlappingevents

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/two-best-non-overlapping-events

func maxTwoEvents(eventsArray [][]int) int {
	n := len(eventsArray)
	events := make([]Event, n)
	for i, e := range eventsArray {
		events[i].startTime, events[i].endTime, events[i].value = e[0], e[1], e[2]
	}
	slices.SortFunc(events, func(a, b Event) int {
		return a.startTime - b.startTime
	})

	// maxValues[i] = j means the maximum value of events[i:n]
	maxValues := make([]int, n)
	maxValues[n-1] = events[n-1].value
	maxPairValue := 0
	for i := n - 2; i >= 0; i-- {
		maxValues[i] = max(events[i].value, maxValues[i+1])
		// find the first event from events[i+1:n) that can start after events[i].
		index, _ := sort.Find(n-i-1, func(j int) int {
			return events[i].endTime + 1 - events[i+1+j].startTime
		})
		if index == n-i-1 {
			continue
		}
		maxPairValue = max(maxPairValue, events[i].value+maxValues[i+1+index])
	}
	return max(maxPairValue, maxValues[0])
}

type Event struct {
	startTime int
	endTime   int
	value     int
}
