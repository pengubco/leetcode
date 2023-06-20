package insert_delete_getrandom_o1

import "math/rand"

/*
In order to use getRandom, we need to store all numbers in an array. But how dow e insert and remove from an array
in O(1)? We insert and remove at the end.
If we need to remove a[i], we switch a[i] and a[n-1] and then remove a[n-1].
*/

type RandomizedSet struct {
	indexes map[int]int
	numbers []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexes: make(map[int]int),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.indexes[val]; ok {
		return false
	}
	n := len(s.numbers)
	s.numbers = append(s.numbers, val)
	s.indexes[val] = n
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	idx, ok := s.indexes[val]
	if !ok {
		return false
	}
	delete(s.indexes, val)
	n := len(s.numbers)
	if idx != n-1 {
		s.numbers[idx] = s.numbers[n-1]
		s.indexes[s.numbers[n-1]] = idx
	}
	s.numbers = s.numbers[:n-1]
	return true
}

func (s *RandomizedSet) GetRandom() int {
	n := len(s.numbers)
	return s.numbers[rand.Intn(n)]
}
