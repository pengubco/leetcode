https://leetcode.com/problems/min-stack/description/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

1. MinStack() initializes the stack object.
1. void push(int val) pushes the element val onto the stack.
1. void pop() removes the element on the top of the stack.
1. int top() gets the top element of the stack.
1. int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

### Example 1:
```text
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]
```

Explanation
1. MinStack minStack = new MinStack();
1. minStack.push(-2);
1. minStack.push(0);
1. minStack.push(-3);
1. minStack.getMin(); // return -3
1. minStack.pop();
1. minStack.top();    // return 0
1. minStack.getMin(); // return -2


## Constraints:

1. `-2^31 <= val <= 2^31 - 1`
1. Methods pop, top and getMin operations will always be called on non-empty stacks.
1. At most `3 * 10^4` calls will be made to push, pop, top, and getMin.
