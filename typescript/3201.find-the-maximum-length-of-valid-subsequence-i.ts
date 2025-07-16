/*
 * @lc app=leetcode.cn id=3201 lang=typescript
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

// @lcpr-template-end
// @lc code=start
function maximumLength(nums: number[]): number {
  const len = nums.length;
  let res = 1;
  const conditions = [
    [0, 0],
    [0, 1],
    [1, 0],
    [1, 1],
  ];

  for (const condition of conditions) {
    let count = 0;
    for (let i = 0; i < len; i++) {
      if (nums[i] % 2 === condition[count % 2]) {
        count++;
      }
    }
    res = Math.max(res, count);
  }

  return res;
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
