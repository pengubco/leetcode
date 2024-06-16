// https://leetcode.com/problems/maximum-gap/

1. Because of the size of input, logN is a constant, log_2(10^5)=17. So technically, using quick sort works.
2. The input is all integers with at most 9 digits, so we can use Radix sort. That brings down the constant to 9.

The two approaches above sort the input. Is there a way that does not sort the input and can get us the answer in two scans
of the data?
Yes. If we can find the following observations.
1. If there an oracle that tells us the max-gap is larger than X, then, we can bucket nums into buckets of range: [min, min+X], [min+X, min+2X],  ....
Because of the max-gap is larger than X, the max-gap cannot come from two numbers from the same bucket. The max-gap must be from two
non-empty adjacent buckets. Assume the two buckets are b[i] and b[j]. The max-gap is min(b[j]) - max(b[i]).

2. So how do we find X? Notice that all we need from X is that just a lower-bound of the max-gap. Assume there are n numbers,
with min and max. One lower-bound is (max-min)/(n-1). We can prove it by contradiction. If the max-gap is smaller than `(max-min)/(n-1)`,
then the total sum of gaps is smaller than (n-1)*(max-min)/(n-1) = max-min, which is the sum of gaps.

So, there we have the X. (max-min)/(n-1). We'll break nums into at most (n-1) buckets. You may think it's weird bucketing n numbers
into (n-1) buckets. What information does the bucketing add? The bucketing let us avoid comparing numbers within the same bucket and
avoid comparing numbers across non-empty non-adjacent buckets.