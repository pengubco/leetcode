
It is attempting to compare the center of a matrix with the target, and decide which 
quadrant to search. However, the time complexity is O(log_{4/3}(N^2))

```text
|A|B|
| x |
|C|D|
```

For example, if x >= target, we can only exclude D. We need to search A, B and C. 

Numbers are increasing from top to bottom and left to right. How may we "concatenate" 
the two increasing sequences to get a longer increasing sequence? 
From top -> bottom -> right.
So we can start from the bottom-left corner. If target is larger, move right. 
Otherwise, move up.