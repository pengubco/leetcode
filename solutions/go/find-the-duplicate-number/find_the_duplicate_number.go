package findtheduplicatenumber

// https://leetcode.com/problems/find-the-duplicate-number/description/

// How to turn the array to a graph?
// The easiest way to construct a graph G is that treat each integer [1,n] as node, and there is directed edge from nums[i] to nums[i+1].
// However, constructing such a graph takes linear space.
// A different way to construct graph G' is forming edge as from i to nums[i].
// Because there is one repeated number, there must be a cycle in G'.
// Proof. Say nums[i] = nums[j] = k.
// i -> nums[i] ( k ) -> nums[k] -> .... -> j -> nums[j] (k)
// why it must go back to j? Becausd there are only n nodes, so we cannot have a path of n+1 without repeating a node.
// I think this is called "group theory" or something.
//
// cycle length: C
// length leads to cycle: P
// meeting point to cycle start: M
//
// say it takes t for slow and fast to meet. then t = P+(C-M)
// the fast traveled P+k*C-M. thus 2*(P+C-M) = P+k*C-M
// 2P + 2C - 2M = P + k*C -M
// P = (k-2)*C + M
// We know that k>=2 because the k=1 for the slow, and fast must have larger k than the slow.
// Therefore, P = k'*C +M, where k'>=0.
// If have another slow' travel from the beginning, then after slow' reaches the cycle start,
// slow must have traveled k'*C + M, which means slow is back at the cycle start.

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}
	return slow
}
