
https://leetcode.com/problems/ugly-number-ii

1. Each ugly number can be calculated from a smaller ugly number by multiply by 2, or 3 or 5.
2. Each ugly number can generate 3 new ugly numbers, by multiply by 2, or 3 or 5.
3. Two different ugly numbers may generate duplicate ugly numbers. For example, 2 * 3 = 3 * 2.
4. To generate the next smallest ugly number, we find the smallest ugly number that has not been 
   multiplied by 2 since it's been created. Do the same for 3 and 5. 
5. How do know which ugly number to use in step 4? Keep numbers generated in an array. 
   The index of ugly-number-to-be-multiplied-by-2 starts with 0. Then every time we multiply the 
	 ugly number by 2, we move the index right. 