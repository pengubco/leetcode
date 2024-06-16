#include <vector>
#include "iostream"

using namespace std;

class Solution {
public:
    int maxProfit(vector<int> &prices) {
        size_t n = prices.size();
        vector<int> firstTrade(n);
        firstTrade[0] = 0;
        int minValue = prices[0];
        for (size_t i = 1; i < n; i++) {
            minValue = min(minValue, prices[i]);
            firstTrade[i] = prices[i] - minValue;
        }
        vector<int> secondTrade(n);
        secondTrade[n - 1] = 0;
        int maxValue = prices[n-1];
        for (int i = n - 2; i >= 0; i--) {
            maxValue = max(maxValue, prices[i]);
            secondTrade[i] = maxValue - prices[i];
        }

        vector<int> maxFirstTrade(n), maxSecondTrade(n);
        maxFirstTrade[0] = firstTrade[0];
        for (int i = 1; i<n; i++) {
            maxFirstTrade[i] = max(maxFirstTrade[i-1], firstTrade[i]);
        }
        maxSecondTrade[n-1] = secondTrade[n-1];
        for (int i = n-2; i>=0; i--) {
            maxSecondTrade[i] = max(maxSecondTrade[i+1], secondTrade[i]);
        }

        int maxProfit = 0;
        for (int i = 0; i < n; i++) {
            maxProfit = max(maxProfit, maxFirstTrade[i] + maxSecondTrade[i]);
        }
        return maxProfit;
    }
};


int main() {
    Solution s;
    vector<int> prices{3, 3, 5, 0, 0, 3, 1, 4};
    cout << s.maxProfit(prices) << endl;

    prices = {1, 2, 3, 4, 5};
    cout << s.maxProfit(prices) << endl;

    prices = {7, 6, 4, 3, 1};
    cout << s.maxProfit(prices) << endl;

    prices = {1, 2, 4, 2, 5, 7, 2, 4, 9, 0};
    cout << s.maxProfit(prices) << endl;
}

