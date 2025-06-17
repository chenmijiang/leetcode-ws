/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30204
 *
 * [27] 移除元素
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/
