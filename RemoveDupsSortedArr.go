/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

package main

import (
	"log"
)

func main() {
	tests := [][]int{{1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}

	for _, test := range tests {
		log.Printf("removeDuplicates(%v) == %d\n", test, removeDuplicates(test))
	}
}

func removeDuplicates(nums []int) int {
	res := len(nums)
	arrayEnd := len(nums)

	for i := 0; i < len(nums); i++ {
		// check if current element is same as prev
		if ((i - 1) >= 0) && (nums[i] == nums[i-1]) {
			// move everything above current element one down
			k := i
			for k < (arrayEnd - 2) {
				nums[k] = nums[k+1]
				k++
			}

			arrayEnd--
			res--
			i--
		}
	}

	return res
}
