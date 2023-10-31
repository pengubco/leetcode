
// https://leetcode.com/problems/reverse-nodes-in-k-group/

Given the head of a linked list, reverse the nodes of the list `k` at a time, and return the modified list.

`k` is a positive integer and is less than or equal to the length of the linked list. 
If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

### Example 1:
![Example 1](./resources/example1.jpg)
```text
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

### Example 2:
![Example 2](./resources/example2.jpg)
```text
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```


## Constraints:
1. The number of nodes in the list is n.
1. `1 <= k <= n <= 5000`
1. `0 <= Node.val <= 1000`

### Follow-up: Can you solve the problem in O(1) extra memory space?
