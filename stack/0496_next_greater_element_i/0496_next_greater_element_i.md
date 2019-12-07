496. Next Greater Element I


Easy


You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2. Find all the next greater numbers for nums1's elements in the corresponding places of nums2.

The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.

Example 1:
```
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
    For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
```

Example 2:

```
Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
    For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
```

Note:   
All elements in nums1 and nums2 are unique.   
The length of both nums1 and nums2 would not exceed 1000.   

## 方法


```go
func nextGreaterElement(findNums []int, nums []int) []int {
    indexOf := make(map[int]int)
	for i, n := range nums {
		indexOf[n] = i
	}

	res := make([]int, len(findNums))
	for i, n := range findNums {
		res[i] = -1
		for j := indexOf[n] + 1; j < len(nums); j++ {
			if n < nums[j] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}
```


```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	res, reocrd, flag := []int{}, map[int]int{}, false
	for i, v := range nums2 {
		reocrd[v] = i
	}
	for i := 0; i < len(nums1); i++ {
		flag = false
		for j := reocrd[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				flag = true
				break
			}
		}
		if flag == false {
			res = append(res, -1)
		}
	}
	return res
}
```


```python
class Solution(object):
    def nextGreaterElement(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[int]
        """
        st, d = [], {}
        for n in nums2:
            while st and st[-1] < n:
                d[st.pop()] = n
            st.append(n)
        return [d.get(x, -1) for x in nums1]
```