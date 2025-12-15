/*
 * @lc app=leetcode.cn id=2110 lang=typescript
 * @lcpr version=30204
 *
 * [2110] 股票平滑下跌阶段的数目
 *
 * https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/description/
 *
 * algorithms
 * Medium (55.50%)
 * Likes:    55
 * Dislikes: 0
 * Total Accepted:    22.4K
 * Total Submissions: 37.3K
 * Testcase Example:  '[3,2,1,4]'
 *
 * 给你一个整数数组 prices ，表示一支股票的历史每日股价，其中 prices[i] 是这支股票第 i 天的价格。
 *
 * 一个 平滑下降的阶段 定义为：对于 连续一天或者多天 ，每日股价都比 前一日股价恰好少 1 ，这个阶段第一天的股价没有限制。
 *
 * 请你返回 平滑下降阶段 的数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：prices = [3,2,1,4]
 * 输出：7
 * 解释：总共有 7 个平滑下降阶段：
 * [3], [2], [1], [4], [3,2], [2,1] 和 [3,2,1]
 * 注意，仅一天按照定义也是平滑下降阶段。
 *
 *
 * 示例 2：
 *
 * 输入：prices = [8,6,7,7]
 * 输出：4
 * 解释：总共有 4 个连续平滑下降阶段：[8], [6], [7] 和 [7]
 * 由于 8 - 6 ≠ 1 ，所以 [8,6] 不是平滑下降阶段。
 *
 *
 * 示例 3：
 *
 * 输入：prices = [1]
 * 输出：1
 * 解释：总共有 1 个平滑下降阶段：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= prices.length <= 10^5
 * 1 <= prices[i] <= 10^5
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function getDescentPeriods(prices: number[]): number {
  // 滑动窗口
  let left = 0;
  let right = 0;
  let n = prices.length;
  let res = 0;

  while (right < n) {
    if (right > 0 && prices[right - 1] - prices[right] === 1) {
      // 符合平滑下降
      right++;
    } else {
      // 不符合平滑下降，计算当前窗口内的平滑下降阶段数目
      let length = right - left;
      res += (length * (length + 1)) / 2;
      left = right;
      right++;
    }
  }
  // 处理最后一个窗口
  let length = right - left;
  res += (length * (length + 1)) / 2;

  return res;
}
// @lc code=end

/*
// @lcpr case=start
// [3,2,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [8,6,7,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

 */
