#include <vector>
#include "iostream"

// https://leetcode.com/problems/kth-largest-element-in-an-array/
class Solution {
public:
    int findKthLargest(std::vector<int>& nums, int k) {
        int low = 0, high = nums.size()-1;
        k = nums.size() - k + 1;
        while (true) {
            int m = this->partition(nums, low, high);
            int count = m-low+1;
            if (count==k) {
                return *std::max_element(std::next(nums.begin(), low), std::next(nums.begin(), m+1));
            } else if (count<k) {
                low = m + 1;
                k -= count;
            } else {
                high = m;
            }
        }
    }

    // partition nums[low, high] to a left part and right part such that,
    // all numbers in the left are less than or equal to numbers in the right part.
    // return the index of the last element of the left group.
    int partition(std::vector<int>& nums, int low, int high) {
        int pivot_index = std::rand() % (high - low + 1) + low;
        std::swap(nums[low], nums[pivot_index]);
        int pivot = nums[low];
        // invariant. nums[low, i] <= pivot <= nums [j, high]
        int i = low-1, j=high+1;
        while(true) {
            do {
                i++;
            } while (nums[i] < pivot);
            do {
                j--;
            } while (nums[j] > pivot);
            if (i >= j) {
                return j;
            }
            std::swap(nums[i], nums[j]);
        }
    }
};

int main() {
    Solution s;
    std::vector<int> nums = {3,2,1,5,6,4};
    std::cout << s.findKthLargest(nums, 2) << std::endl;
    nums = {3,2,3,1,2,4,5,5,6};
    std::cout << s.findKthLargest(nums, 4) << std::endl;
    nums = {1};
    std::cout << s.findKthLargest(nums, 1) << std::endl;

    nums = {7,6,5,4,3,2,1};
    std::cout << s.findKthLargest(nums, 5) << std::endl;

    return 0;
}
