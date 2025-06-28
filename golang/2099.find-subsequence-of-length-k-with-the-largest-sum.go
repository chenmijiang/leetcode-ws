/*
 * @lc app=leetcode.cn id=2099 lang=golang
 * @lcpr version=30204
 *
 * [2099] 找到和最大的长度为 K 的子序列
 *
 * https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum/description/
 *
 * algorithms
 * Easy (49.96%)
 * Likes:    66
 * Dislikes: 0
 * Total Accepted:    21.8K
 * Total Submissions: 38.4K
 * Testcase Example:  '[2,1,3,3]\n2'
 *
 * 给你一个整数数组 nums 和一个整数 k 。你需要找到 nums 中长度为 k 的 子序列 ，且这个子序列的 和最大 。
 *
 * 请你返回 任意 一个长度为 k 的整数子序列。
 *
 * 子序列 定义为从一个数组里删除一些元素后，不改变剩下元素的顺序得到的数组。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,1,3,3], k = 2
 * 输出：[3,3]
 * 解释：
 * 子序列有最大和：3 + 3 = 6 。
 *
 * 示例 2：
 *
 * 输入：nums = [-1,-2,3,4], k = 3
 * 输出：[-1,3,4]
 * 解释：
 * 子序列有最大和：-1 + 3 + 4 = 6 。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [3,4,3,3], k = 2
 * 输出：[3,4]
 * 解释：
 * 子序列有最大和：3 + 4 = 7 。
 * 另一个可行的子序列为 [4, 3] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 1000
 * -10^5 <= nums[i] <= 10^5
 * 1 <= k <= nums.length
 *
 *
 */

// @lcpr-template-start
package main

import "sort"

// @lcpr-template-end
// @lc code=start
func maxSubsequence(nums []int, k int) []int {
	// 思路：创建索引数组，按值排序，取前k个最大值的索引，再按原始索引排序

	indices := make([]int, len(nums))
	for i := range indices {
		indices[i] = i
	}

	// 按值从大到小排序索引
	sort.Slice(indices, func(i, j int) bool {
		if nums[indices[i]] == nums[indices[j]] {
			return indices[i] < indices[j] // 值相等时按索引升序
		}
		return nums[indices[i]] > nums[indices[j]]
	})

	// 取前k个索引
	selectedIndices := indices[:k]

	// 按原始索引排序以保持原有顺序
	sort.Ints(selectedIndices)

	// 构建结果
	result := make([]int, k)
	for i, idx := range selectedIndices {
		result[i] = nums[idx]
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [-1,-2,3,4]\n3\n
// @lcpr case=end

// @lcpr case=start
// [3,4,3,3]\n2\n
// @lcpr case=end

*/
