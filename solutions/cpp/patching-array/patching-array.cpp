#include <vector>
#include "iostream"

// https://leetcode.com/problems/patching-array/
class Solution {
public:
    int minPatches(std::vector<int>& nums, int n) {
        int patches = 0;
        // all integers between [0, m] can be formed.
        int64_t m = 0;
        int used_index = 0;
        for (;m < n;) {
            if (used_index < nums.size() && nums[used_index] <= (m+1)) {
                m += nums[used_index];
                used_index++;
            } else {
                m += (m+1);
                patches++;
            }
        }
        return patches;
    }
};

int main() {
    Solution s;
    std::vector<int> nums = {1,3};
    std::cout << s.minPatches(nums, 6) << std::endl;

    nums = {1, 5, 10};
    std::cout << s.minPatches(nums, 20) << std::endl;

    nums = {1, 2, 2};
    std::cout << s.minPatches(nums, 5) << std::endl;

    nums = {1,2,31,33};
    std::cout << s.minPatches(nums, 2147483647) << std::endl;
}
