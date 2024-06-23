#include <algorithm>
#include <climits>
#include <iostream>
#include <map>
#include <numeric>
#include <vector>

using namespace std;

// https://leetcode.com/problems/minimum-size-subarray-in-infinite-array/description/
class Solution {
public:
    int minSizeSubarray(vector<int> &nums, int target) {
      int64_t totalSum = std::accumulate(nums.begin(), nums.end(), int64_t (0));
      int copies = target / totalSum;
      target %= totalSum;
      int n = nums.size();
      map<int64_t , int> sumToIndex;
      sumToIndex[0] = -1;

      int64_t sum = 0;
      int minSize = INT_MAX;
      for (int i = 0; i < 2 * n; i++) {
        sum += nums[i % n];
        sumToIndex[sum] = i;
        auto iter = sumToIndex.find(sum - target);
        if (iter == sumToIndex.end()) {
          continue;
        }
        minSize = min(minSize, i - iter->second);
      }
      return (minSize == INT_MAX) ? -1 :  minSize + copies * n;
    };
};

int main() {
  Solution s;
  vector<int> nums;
  int target;
  int expected;
  int result;

  nums = {2, 1, 5, 7, 7, 1, 6, 3};
  target = 39;
  expected = 9;
  result = s.minSizeSubarray(nums, target);
  if (result != expected) {
    cout << "wrong " << result << endl;
    return 1;
  }

  nums = {1, 2, 3};
  target = 5;
  expected = 2;
  result = s.minSizeSubarray(nums, target);
  if (result != expected) {
    cout << "wrong " << result << endl;
    return 1;
  }

  nums = {1, 1, 1, 2, 3};
  target = 4;
  expected = 2;
  result = s.minSizeSubarray(nums, target);
  if (result != expected) {
    cout << "wrong " << result << endl;
    return 1;
  }

  nums = {2, 4, 6, 8};
  target = 3;
  expected = -1;
  result = s.minSizeSubarray(nums, target);
  if (result != expected) {
    cout << "wrong " << result << endl;
    return 1;
  }
}

