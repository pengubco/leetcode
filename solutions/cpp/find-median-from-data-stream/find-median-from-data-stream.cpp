#include <queue>
using namespace std;

// https://leetcode.com/problems/find-median-from-data-stream/description/

class MedianFinder {
public:
    MedianFinder() {
    }

    void addNum(int num) {
        this->max_queue.push(num);
        this->total++;
        while(this->min_queue.size() < this->total / 2) {
            this->min_queue.push(this->max_queue.top());
            this->max_queue.pop();
        }
        if (this->min_queue.empty()) {
            return;
        }
        while(this->min_queue.top() < this->max_queue.top()) {
            auto m = this->max_queue.top();
            auto n = this->min_queue.top();
            this->max_queue.pop();
            this->min_queue.pop();
            this->max_queue.push(n);
            this->min_queue.push(m);
        }
    }

    double findMedian() {
        if (this->total % 2 == 0) {
            return double(this->max_queue.top() + this->min_queue.top()) / 2;
        }
        return double(this->max_queue.top());
    }

    priority_queue<int> max_queue;
    priority_queue<int, vector<int>, greater<int>> min_queue;
    size_t total = 0;
};

int main() {
    MedianFinder *medianFinder = new MedianFinder();
    medianFinder->addNum(1);    // arr = [1]
    medianFinder->addNum(2);    // arr = [1, 2]
    printf("%f\n", medianFinder->findMedian()); // return 1.5 (i.e., (1 + 2) / 2)
    medianFinder->addNum(3);    // arr[1, 2, 3]
    printf("%f\n",medianFinder->findMedian()); // return 2.0
    return 0;
}
