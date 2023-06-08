
// https://leetcode.com/problems/max-points-on-a-line/description/

Given an array of points where `points[i] = [x_i, y_i]` represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.

### Example 1:
```text
Input: points = [[1,1],[2,2],[3,3]]
Output: 3
```
![example1.jpg](./resources/example1.jpg)

### Example 2:
```text
Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
Output: 4
```

![example2.jpg](./resources/example2.jpg)

## Constraints:

1. `1 <= points.length <= 300`
1. `points[i].length == 2`
1. `-10^4 <= x_i, y_i <= 10^4`
1. All the points are unique.