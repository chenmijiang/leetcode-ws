/*
 * @lc app=leetcode.cn id=169 lang=golang
 * @lcpr version=30204
 *
 * [169] 多数元素
 *
 * https://leetcode.cn/problems/majority-element/description/
 *
 * algorithms
 * Easy (66.86%)
 * Likes:    2501
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 1.9M
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,2,3]
 * 输出：3
 *
 * 示例 2：
 *
 * 输入：nums = [2,2,1,1,1,2,2]
 * 输出：2
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 *
 * 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func majorityElement(nums []int) int {
	// 候选众数
	candidate := nums[0]
	// 计数器
	count := 1

	// Boyer-Moore 投票算法
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			// 更换候选众数
			candidate = nums[i]
			count = 1
		} else if nums[i] == candidate {
			// 遇到相同的数，计数器加1
			count++
		} else {
			// 遇到不同的数，计数器减1
			count--
		}
	}

	return candidate
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,1,1,2,2]\n
// @lcpr case=end

*/
