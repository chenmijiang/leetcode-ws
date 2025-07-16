/*
 * @lc app=leetcode.cn id=3201 lang=golang
 * @lcpr version=30204
 *
 * [3201] 找出有效子序列的最大长度 I
 *
 * https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-i/description/
 *
 * algorithms
 * Medium (38.74%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    15.5K
 * Total Submissions: 28.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个整数数组 nums。
 *
 * nums 的子序列 sub 的长度为 x ，如果其满足以下条件，则称其为 有效子序列：
 *
 *
 * (sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x
 * - 1]) % 2
 *
 *
 * 返回 nums 的 最长的有效子序列 的长度。
 *
 * 一个 子序列 指的是从原数组中删除一些元素（也可以不删除任何元素），剩余元素保持原来顺序组成的新数组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入： nums = [1,2,3,4]
 *
 * 输出： 4
 *
 * 解释：
 *
 * 最长的有效子序列是 [1, 2, 3, 4]。
 *
 *
 * 示例 2：
 *
 *
 * 输入： nums = [1,2,1,1,2,1,2]
 *
 * 输出： 6
 *
 * 解释：
 *
 * 最长的有效子序列是 [1, 2, 1, 2, 1, 2]。
 *
 *
 * 示例 3：
 *
 *
 * 输入： nums = [1,3]
 *
 * 输出： 2
 *
 * 解释：
 *
 * 最长的有效子序列是 [1, 3]。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 2 * 10^5
 * 1 <= nums[i] <= 10^7
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func maximumLength(nums []int) int {
	// 一次遍历，位运算

	// 全奇
	var len1 = 1 - nums[0]&1
	// 全偶
	var len2 = nums[0] & 1
	// 奇偶相见
	var len3 = 1

	for i := 1; i < len(nums); i++ {
		len1 += 1 - nums[i]&1
		len2 += nums[i] & 1
		len3 += (nums[i] ^ nums[i-1]) & 1
	}

	return max(len1, max(len2, len3))
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,1,2,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5]\n
// @lcpr case=end

*/
