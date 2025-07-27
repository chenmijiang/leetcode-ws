/*
 * @lc app=leetcode.cn id=2210 lang=golang
 * @lcpr version=30204
 *
 * [2210] 统计数组中峰和谷的数量
 *
 * https://leetcode.cn/problems/count-hills-and-valleys-in-an-array/description/
 *
 * algorithms
 * Easy (59.62%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    20.1K
 * Total Submissions: 31.7K
 * Testcase Example:  '[2,4,1,1,6,5]'
 *
 * 给你一个下标从 0 开始的整数数组 nums 。如果两侧距 i 最近的不相等邻居的值均小于 nums[i] ，则下标 i 是 nums
 * 中，某个峰的一部分。类似地，如果两侧距 i 最近的不相等邻居的值均大于 nums[i] ，则下标 i 是 nums 中某个谷的一部分。对于相邻下标 i
 * 和 j ，如果 nums[i] == nums[j] ， 则认为这两下标属于 同一个 峰或谷。
 *
 * 注意，要使某个下标所做峰或谷的一部分，那么它左右两侧必须 都 存在不相等邻居。
 *
 * 返回 nums 中峰和谷的数量。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,4,1,1,6,5]
 * 输出：3
 * 解释：
 * 在下标 0 ：由于 2 的左侧不存在不相等邻居，所以下标 0 既不是峰也不是谷。
 * 在下标 1 ：4 的最近不相等邻居是 2 和 1 。由于 4 > 2 且 4 > 1 ，下标 1 是一个峰。
 * 在下标 2 ：1 的最近不相等邻居是 4 和 6 。由于 1 < 4 且 1 < 6 ，下标 2 是一个谷。
 * 在下标 3 ：1 的最近不相等邻居是 4 和 6 。由于 1 < 4 且 1 < 6 ，下标 3 符合谷的定义，但需要注意它和下标 2
 * 是同一个谷的一部分。
 * 在下标 4 ：6 的最近不相等邻居是 1 和 5 。由于 6 > 1 且 6 > 5 ，下标 4 是一个峰。
 * 在下标 5 ：由于 5 的右侧不存在不相等邻居，所以下标 5 既不是峰也不是谷。
 * 共有 3 个峰和谷，所以返回 3 。
 *
 * 示例 2：
 *
 * 输入：nums = [6,6,5,5,4,1]
 * 输出：0
 * 解释：
 * 在下标 0 ：由于 6 的左侧不存在不相等邻居，所以下标 0 既不是峰也不是谷。
 * 在下标 1 ：由于 6 的左侧不存在不相等邻居，所以下标 1 既不是峰也不是谷。
 * 在下标 2 ：5 的最近不相等邻居是 6 和 4 。由于 5 < 6 且 5 > 4 ，下标 2 既不是峰也不是谷。
 * 在下标 3 ：5 的最近不相等邻居是 6 和 4 。由于 5 < 6 且 5 > 4 ，下标 3 既不是峰也不是谷。
 * 在下标 4 ：4 的最近不相等邻居是 5 和 1 。由于 4 < 5 且 4 > 1 ，下标 4 既不是峰也不是谷。
 * 在下标 5 ：由于 1 的右侧不存在不相等邻居，所以下标 5 既不是峰也不是谷。
 * 共有 0 个峰和谷，所以返回 0 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func countHillValley(nums []int) int {
	var ans, preStatus int

	for i := 0; i < len(nums)-1; i++ {
		curStatus := sign(nums[i] - nums[i+1]) // 1: 峰, -1: 谷, 0: 平

		if curStatus == 0 {
			continue
		}

		if preStatus == -curStatus {
			ans++
		}

		preStatus = curStatus
	}

	return ans
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,1,1,6,5]\n
// @lcpr case=end

// @lcpr case=start
// [6,6,5,5,4,1]\n
// @lcpr case=end

*/
