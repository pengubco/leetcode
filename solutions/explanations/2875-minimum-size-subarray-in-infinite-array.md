Any subarray of the infinite array is comprised of three parts.
1. A suffix of the original array.
2. A prefix of the original array.
3. K copies of the original array.
Each part can be empty.

We can further simplifies by combining part 1 and part 2 to the following.
A subarray of two copies of the original array.
a[i:j] + k * total_sum = target. `a` is the original array concatenated to itself.

If we know `k`, then the problem is simple: find a minimum subarray with a fixed sum. 

Can we say k is as large as possible? that is, `k=target/total_sum`? 

Yes. This is because all numbers are positive, the shortest subarray whose sum is 
total_sum must include each number once and no more than once. 
To make it more clear, by repeating an array, A, to itself multiple times. 
A[i] and A[j] where A[j] is the i-th number in a repeating array, then 
`j-i = k*n`, where `k>1`.
