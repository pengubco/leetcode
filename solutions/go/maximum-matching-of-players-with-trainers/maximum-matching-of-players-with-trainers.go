package maximummatchingofplayerswithtrainers

import (
	"sort"
)

// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/

func matchPlayersAndTrainers(players []int, trainers []int) int {
	result := 0
	sort.Ints(players)
	sort.Ints(trainers)
	n, m := len(players), len(trainers)
	for i, j := 0, 0; i < n && j < m; {
		if players[i] <= trainers[j] {
			i++
			j++
			result++
			continue
		}
		j++
	}
	return result
}
