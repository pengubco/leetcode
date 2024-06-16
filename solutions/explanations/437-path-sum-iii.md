
1. define top-down path as: a path a to b where a is b's ancestor. 
2. any top-down path is a suffix of a top-down path from root. 
3. A suffix path is a top-down path minus a prefix path.
4. A top-down path from root is a path from the DFS at a point in time.
So we can DFS from root, maintain a map, m:
path-sum -> count of top-down paths with the sum.
When visiting a node of v first time, sum up m[targetSum - v].

initial condition: m[0]=1