/*
 * @lc app=leetcode.cn id=2016 lang=golang
 * @lcpr version=30204
 *
 * [2016] 增量元素之间的最大差值
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func maximumDifference(nums []int) int {
	min := nums[0]
	max := -1
	for _, num := range nums {
		if min >= num {
			min = num
			continue
		}
		if num-min > max {
			max = num - min
		}
	}
	return max
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,4]\n
// @lcpr case=end

// @lcpr case=start
// [9,4,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,5,2,10]\n
// @lcpr case=end

*/
