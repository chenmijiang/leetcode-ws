/*
 * @lc app=leetcode.cn id=3487 lang=typescript
 * @lcpr version=30204
 *
 * [3487] 删除后的最大子数组元素和
 *
 * https://leetcode.cn/problems/maximum-unique-subarray-sum-after-deletion/description/
 *
 * algorithms
 * Easy (41.37%)
 * Likes:    19
 * Dislikes: 0
 * Total Accepted:    16K
 * Total Submissions: 32.8K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给你一个整数数组 nums 。
 * 
 * 你可以从数组 nums 中删除任意数量的元素，但不能将其变为 空 数组。执行删除操作后，选出 nums 中满足下述条件的一个子数组：
 * 
 * 
 * 子数组中的所有元素 互不相同 。
 * 最大化 子数组的元素和。
 * 
 * 
 * 返回子数组的 最大元素和 。
 * 子数组 是数组的一个连续、非空 的元素序列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4,5]
 * 
 * 输出：15
 * 
 * 解释：
 * 
 * 不删除任何元素，选中整个数组得到最大元素和。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,1,0,1,1]
 * 
 * 输出：1
 * 
 * 解释：
 * 
 * 删除元素 nums[0] == 1、nums[1] == 1、nums[2] == 0 和 nums[3] == 1 。选中整个数组 [1]
 * 得到最大元素和。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1,2,-1,-2,1,0,-1]
 * 
 * 输出：3
 * 
 * 解释：
 * 
 * 删除元素 nums[2] == -1 和 nums[3] == -2 ，从 [1, 2, 1, 0, -1] 中选中子数组 [2, 1]
 * 以获得最大元素和。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * -100 <= nums[i] <= 100
 * 
 * 
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function maxSum(nums: number[]): number {
  // 大于 0 去重set，如果长度 为 0，则取 nums 中 最大值，反之求和
  const set = new Set<number>(nums.filter(item => item > 0));
  if (set.size === 0) {
    return Math.max(...nums);
  }
  return Array.from(set).reduce((acc, item) => acc + item, 0);
};
// @lc code=end



/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,-1,-2,1,0,-1]\n
// @lcpr case=end

 */

