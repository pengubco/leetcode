
This is the minimum vertex cover problem. It is NPC on graphs. But for trees, it 
can be solved in O(N). 

1. No matter which node is the root, the size of minimum vertex cover does not change. 
   This is because a vertex cover covers all edges. The set of all edges is the same 
	 no matter which node is picked as root. Let T1 and T2 be two trees using different 
	 root nodes, a cover for T1 is also a cover for T2, vice versa. So the minimum 
	 vertex cover must have the same size. 
2. Fix the root. For any leaf, it is always better to pick its parent. Remove the leaf, 
   its parent and all other nodes covered by its parent. Repeat till there is no leaf.
3. In implementation, Post-order DFS gives us "leaf" after "deletion". Avoid changing node's 
   Left or Right to nil for "deletion". 

