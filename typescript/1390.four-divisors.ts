/*
 * @lc app=leetcode.cn id=1390 lang=typescript
 * @lcpr version=30204
 *
 * [1390] 四因数
 *
 * https://leetcode.cn/problems/four-divisors/description/
 *
 * algorithms
 * Medium (41.32%)
 * Likes:    47
 * Dislikes: 0
 * Total Accepted:    16K
 * Total Submissions: 34.8K
 * Testcase Example:  '[21,4,7]'
 *
 * 给你一个整数数组 nums，请你返回该数组中恰有四个因数的这些整数的各因数之和。如果数组中不存在满足题意的整数，则返回 0 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [21,4,7]
 * 输出：32
 * 解释：
 * 21 有 4 个因数：1, 3, 7, 21
 * 4 有 3 个因数：1, 2, 4
 * 7 有 2 个因数：1, 7
 * 答案仅为 21 的所有因数的和。
 *
 *
 * 示例 2:
 *
 * 输入: nums = [21,21]
 * 输出: 64
 *
 *
 * 示例 3:
 *
 * 输入: nums = [1,2,3,4,5]
 * 输出: 0
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * 1 <= nums[i] <= 10^5
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function sumFourDivisors(nums: number[]): number {
  let totalSum = 0;

  for (const num of nums) {
    let count = 0;
    let currentSum = 0;

    const limit = Math.sqrt(num);
    for (let i = 0; i <= limit; i++) {
      if (num % i !== 0) {
        continue;
      }

      const complement = num / i;
      if (complement === i) {
        count++;
        currentSum += i;
      } else {
        count += 2;
        currentSum = currentSum + i + complement;
      }

      if (count > 4) {
        break;
      }
    }
    if (count === 4) {
      totalSum += currentSum;
    }
  }

  return totalSum;
}
// @lc code=end

/*
// @lcpr case=start
// [21,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [21,21]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

 */
