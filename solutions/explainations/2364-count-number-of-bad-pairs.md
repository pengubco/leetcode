
// https://leetcode.com/problems/count-number-of-bad-pairs

bad pairs are all pairs minus valid pairs.
A pair is valid if `j < i and a[i]-a[j] = i-j`. i.e., a[i]-i == a[j]-j. With this observation, we can bucket numbers in a[i]-i.
1. If there are n numbers in the same bucket, then the number of valid pairs from the bucket is n*(n-1)/2.
2. Only numbers from the same bucket can compose a valid pair.