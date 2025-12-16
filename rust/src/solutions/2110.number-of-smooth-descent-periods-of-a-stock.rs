/*
 * @lc app=leetcode.cn id=2110 lang=rust
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

// @lc code=start
use crate::Solution;

impl Solution {
    pub fn get_descent_periods(prices: Vec<i32>) -> i64 {
        let mut result = 0i64;
        let mut count = 1i64;
        for i in 1..prices.len() {
            if prices[i - 1] - prices[i] == 1 {
                count += 1;
            } else {
                result += count * (count + 1) / 2;
                count = 1;
            }
        }
        result += count * (count + 1) / 2;
        result
    }
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example_1() {
        // 示例1：[3,2,1,4] -> 7
        // 平滑下降阶段：[3], [2], [1], [4], [3,2], [2,1], [3,2,1]
        let prices = vec![3, 2, 1, 4];
        assert_eq!(Solution::get_descent_periods(prices), 7);
    }

    #[test]
    fn test_example_2() {
        // 示例2：[8,6,7,7] -> 4
        // 平滑下降阶段：[8], [6], [7], [7]
        // 8-6=2 ≠ 1，所以[8,6]不是平滑下降阶段
        let prices = vec![8, 6, 7, 7];
        assert_eq!(Solution::get_descent_periods(prices), 4);
    }

    #[test]
    fn test_example_3() {
        // 示例3：[1] -> 1
        // 单个元素也是一个平滑下降阶段
        let prices = vec![1];
        assert_eq!(Solution::get_descent_periods(prices), 1);
    }

    #[test]
    fn test_all_smooth_descent() {
        // 边界情况：完全递减序列 [5,4,3,2,1]
        // 平滑下降阶段数 = n*(n+1)/2 = 5*6/2 = 15
        let prices = vec![5, 4, 3, 2, 1];
        assert_eq!(Solution::get_descent_periods(prices), 15);
    }

    #[test]
    fn test_no_consecutive_descent() {
        // 边界情况：没有连续下降 [1,3,5,7]
        // 每个元素单独构成一个平滑下降阶段
        let prices = vec![1, 3, 5, 7];
        assert_eq!(Solution::get_descent_periods(prices), 4);
    }

    #[test]
    fn test_two_elements_descent() {
        // 边界情况：两个元素恰好差1 [2,1]
        // 平滑下降阶段：[2], [1], [2,1] = 3
        let prices = vec![2, 1];
        assert_eq!(Solution::get_descent_periods(prices), 3);
    }

    #[test]
    fn test_two_elements_no_descent() {
        // 边界情况：两个元素差值不为1 [5,3]
        // 平滑下降阶段：[5], [3] = 2
        let prices = vec![5, 3];
        assert_eq!(Solution::get_descent_periods(prices), 2);
    }

    #[test]
    fn test_multiple_descent_segments() {
        // 边界情况：多个下降段 [5,4,3,6,5,4]
        // 第一段 [5,4,3]: 1+2+3 = 6 个子数组
        // 第二段 [6,5,4]: 1+2+3 = 6 个子数组
        // 总共 = 12
        let prices = vec![5, 4, 3, 6, 5, 4];
        assert_eq!(Solution::get_descent_periods(prices), 12);
    }

    #[test]
    fn test_same_values() {
        // 边界情况：所有值相同 [3,3,3]
        // 每个元素单独构成一个平滑下降阶段 = 3
        let prices = vec![3, 3, 3];
        assert_eq!(Solution::get_descent_periods(prices), 3);
    }
}
