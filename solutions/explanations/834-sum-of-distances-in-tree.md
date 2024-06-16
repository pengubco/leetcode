
1. Assume the tree is rooted. i is j's parent. Then 
  distanceSum[j] = distanceSum[i] - subtreeSize[j] + (n - subtreeSize[j]) 

n is the size of the tree.
subtreeSize[j] is the size of the subtree rooted at j. 

So the solution is
1. DFS to make the tree rooted at 0. The calculate distanceSum[0]
2. Then DFS again to calculate distanceSum for all other nodes. 