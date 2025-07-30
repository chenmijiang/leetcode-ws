/*
 * @lc app=leetcode.cn id=2419 lang=golang
 * @lcpr version=30204
 *
 * [2419] 按位与最大的最长子数组
 *
 * https://leetcode.cn/problems/longest-subarray-with-maximum-bitwise-and/description/
 *
 * algorithms
 * Medium (45.23%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    21.7K
 * Total Submissions: 42.9K
 * Testcase Example:  '[1,2,3,3,2,2]'
 *
 * 给你一个长度为 n 的整数数组 nums 。
 *
 * 考虑 nums 中进行 按位与（bitwise AND）运算得到的值 最大 的 非空 子数组。
 *
 *
 * 换句话说，令 k 是 nums 任意 子数组执行按位与运算所能得到的最大值。那么，只需要考虑那些执行一次按位与运算后等于 k 的子数组。
 *
 *
 * 返回满足要求的 最长 子数组的长度。
 *
 * 数组的按位与就是对数组中的所有数字进行按位与运算。
 *
 * 子数组 是数组中的一个连续元素序列。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,2,3,3,2,2]
 * 输出：2
 * 解释：
 * 子数组按位与运算的最大值是 3 。
 * 能得到此结果的最长子数组是 [3,3]，所以返回 2 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [1,2,3,4]
 * 输出：1
 * 解释：
 * 子数组按位与运算的最大值是 4 。
 * 能得到此结果的最长子数组是 [4]，所以返回 1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func longestSubarray(nums []int) int {
	ans, mv, count := 0, 0, 0

	for _, num := range nums {
		if num > mv {
			mv = num
			ans, count = 1, 1
		} else if num == mv {
			count++
			ans = max(ans, count)
		} else {
			count = 0
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

*/
