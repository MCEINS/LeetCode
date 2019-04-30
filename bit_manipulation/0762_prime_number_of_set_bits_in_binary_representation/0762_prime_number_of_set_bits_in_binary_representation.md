762. Prime Number of Set Bits in Binary Representation

Easy

Given two integers L and R, find the count of numbers in the range [L, R] (inclusive) having a prime number of set bits in their binary representation.

(Recall that the number of set bits an integer has is the number of 1s present when written in binary. For example, 21 written in binary is 10101 which has 3 set bits. Also, 1 is not a prime.)

Example 1:

Input: L = 6, R = 10  
Output: 4  
Explanation:  
6 -> 110 (2 set bits, 2 is prime)  
7 -> 111 (3 set bits, 3 is prime)  
9 -> 1001 (2 set bits , 2 is prime)  
10->1010 (2 set bits , 2 is prime)  
Example 2:  
 
Input: L = 10, R = 15  
Output: 5  
Explanation:  
10 -> 1010 (2 set bits, 2 is prime)  
11 -> 1011 (3 set bits, 3 is prime)  
12 -> 1100 (2 set bits, 2 is prime)  
13 -> 1101 (3 set bits, 3 is prime)  
14 -> 1110 (3 set bits, 3 is prime)  
15 -> 1111 (4 set bits, 4 is not prime)  
Note:  

L, R will be integers L <= R in the range [1, 10^6].
R - L will be at most 10000.

# 方法  
1）R - L 最多 10^6，转换为二进制最多20 位，可以罗列出范围内所有的素数，依次遍历判断计数
2）利用 python 内置函数直接算出二进制 1 的个数，遍历计数



```go
func countPrimeSetBits(L int, R int) int {
	primes := [...]int{2: 1, 3: 1, 5: 1, 7: 1, 11: 1, 13: 1, 17: 1, 19: 1}

	res := 0
	for i := L; i <= R; i++ {
		bits := 0
		for n := i; n > 0; n >>= 1 {
			bits += n & 1
		}
		res += primes[bits]
	}

	return res
}
```


```python
class Solution(object):
    def countPrimeSetBits(self, L, R):
        """
        :type L: int
        :type R: int
        :rtype: int
        """
        primes = {2, 3, 5, 7, 11, 13, 17, 19}
        return sum(map(lambda x: bin(x).count('1') in primes, range(L, R+1)))
        
```