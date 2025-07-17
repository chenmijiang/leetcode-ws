/*
 * @lc app=leetcode.cn id=3202 lang=golang
 * @lcpr version=30204
 *
 * [3202] 找出有效子序列的最大长度 II
 *
 * https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-ii/description/
 *
 * algorithms
 * Medium (46.01%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    14.8K
 * Total Submissions: 25K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个整数数组 nums 和一个 正 整数 k 。
 * nums 的一个 子序列 sub 的长度为 x ，如果其满足以下条件，则称其为 有效子序列 ：
 *
 *
 * (sub[0] + sub[1]) % k == (sub[1] + sub[2]) % k == ... == (sub[x - 2] + sub[x
 * - 1]) % k
 *
 * 返回 nums 的 最长有效子序列 的长度。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,4,5], k = 2
 *
 * 输出：5
 *
 * 解释：
 *
 * 最长有效子序列是 [1, 2, 3, 4, 5] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,4,2,3,1,4], k = 3
 *
 * 输出：4
 *
 * 解释：
 *
 * 最长有效子序列是 [1, 4, 1, 4] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^3
 * 1 <= nums[i] <= 10^7
 * 1 <= k <= 10^3
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func maximumLength(nums []int, k int) int {
	// dp[i][j] 表示：上一个数字是 i，当前数字是 j 时的最长子序列长度
	dp := make([][]int, k)

	for i := range dp {
		dp[i] = make([]int, k)
	}

	res := 0

	// 遍历数组中的每个数字
	for _, num := range nums {
		num %= k // 只关心数字对 k 的余数

		// 尝试把当前数字接在所有可能的前一个数字后面
		for prev := 0; prev < k; prev++ {
			// 关键：dp[prev][num] 表示前一个数是 prev，当前数是 num 的最长长度
			// dp[num][prev] + 1 表示在原来的基础上加上当前数字
			dp[prev][num] = dp[num][prev] + 1
			res = max(res, dp[prev][num])
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,2,3,1,4]\n3\n
// @lcpr case=end

*/
