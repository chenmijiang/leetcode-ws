/*
 * @lc app=leetcode.cn id=1390 lang=rust
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
use crate::Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn sum_four_divisors(nums: Vec<i32>) -> i32 {
        let mut total_sum = 0;

        for &num in &nums {
            let mut count = 0;
            let mut sum = 0;
            let sqrt_n = (num as f64).sqrt() as i32;

            for i in 1..=sqrt_n {
                if num % i != 0 {
                    continue;
                }

                let other_divisor = num / i;

                if other_divisor == i {
                    count += 1;
                    sum += i;
                } else {
                    count += 2;
                    sum += i + other_divisor;
                }

                if count > 4 {
                    break;
                }
            }

            if count == 4 {
                total_sum += sum;
            }
        }

        total_sum
    }
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
