/*
 * @lc app=leetcode.cn id=594 lang=golang
 * @lcpr version=30204
 *
 * [594] 最长和谐子序列
 *
 * https://leetcode.cn/problems/longest-harmonious-subsequence/description/
 *
 * algorithms
 * Easy (56.91%)
 * Likes:    431
 * Dislikes: 0
 * Total Accepted:    114.4K
 * Total Submissions: 196.2K
 * Testcase Example:  '[1,3,2,2,5,2,3,7]'
 *
 * 和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
 *
 * 给你一个整数数组 nums ，请你在所有可能的 子序列 中找到最长的和谐子序列的长度。
 *
 * 数组的 子序列 是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,3,2,2,5,2,3,7]
 *
 * 输出：5
 *
 * 解释：
 *
 * 最长和谐子序列是 [3,2,2,2,3]。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4]
 *
 * 输出：2
 *
 * 解释：
 *
 * 最长和谐子序列是 [1,2]，[2,3] 和 [3,4]，长度都为 2。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,1,1,1]
 *
 * 输出：0
 *
 * 解释：
 *
 * 不存在和谐子序列。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func findLHS(nums []int) (ans int) {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
		if mp[num-1] > 0 {
			ans = max(ans, mp[num]+mp[num-1])
		}
		if mp[num+1] > 0 {
			ans = max(ans, mp[num]+mp[num+1])
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,2,5,2,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n
// @lcpr case=end

*/
