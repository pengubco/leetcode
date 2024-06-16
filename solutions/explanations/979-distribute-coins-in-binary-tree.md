
1. If a node needs a coin, then the coin must come from the closest node. This is because, 
   getting coin from a further node doesn't help the overall solution. In the end, 
	 every node has one coin. So we can consider each node independently. 
2. Repeating getting coin from the nearest node with excessive coin is difficult to implement 
   because there is no parent edge. 
3. Consider any subtree of x nodes and y coins. 
   1. If y>x, then it must contribute y-x moves to its parent. 
   2. If y==x, no move between the subtree and its parent. 
   3. Otherwise, the its parent must contribute x-y moves to it. 
 