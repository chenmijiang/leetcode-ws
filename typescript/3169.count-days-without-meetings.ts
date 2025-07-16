/*
 * @lc app=leetcode.cn id=3169 lang=typescript
 * @lcpr version=30204
 *
 * [3169] 无需开会的工作日
 *
 * https://leetcode.cn/problems/count-days-without-meetings/description/
 *
 * algorithms
 * Medium (39.43%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    18K
 * Total Submissions: 38K
 * Testcase Example:  '10\n[[5,7],[1,3],[9,10]]'
 *
 * 给你一个正整数 days，表示员工可工作的总天数（从第 1 天开始）。另给你一个二维数组 meetings，长度为 n，其中 meetings[i] =
 * [start_i, end_i] 表示第 i 次会议的开始和结束天数（包含首尾）。
 *
 * 返回员工可工作且没有安排会议的天数。
 *
 * 注意：会议时间可能会有重叠。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：days = 10, meetings = [[5,7],[1,3],[9,10]]
 *
 * 输出：2
 *
 * 解释：
 *
 * 第 4 天和第 8 天没有安排会议。
 *
 *
 * 示例 2：
 *
 *
 * 输入：days = 5, meetings = [[2,4],[1,3]]
 *
 * 输出：1
 *
 * 解释：
 *
 * 第 5 天没有安排会议。
 *
 *
 * 示例 3：
 *
 *
 * 输入：days = 6, meetings = [[1,6]]
 *
 * 输出：0
 *
 * 解释：
 *
 * 所有工作日都安排了会议。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= days <= 10^9
 * 1 <= meetings.length <= 10^5
 * meetings[i].length == 2
 * 1 <= meetings[i][0] <= meetings[i][1] <= days
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function countDays(days: number, meetings: number[][]): number {
  meetings.sort((a, b) => a[0] - b[0]);

  let count = 0;
  let end = 1;
  let len = meetings.length;

  // 计算会议之间的空闲天数
  for (let i = 0; i < len; i++) {
    const [startTime, endTime] = meetings[i];
    // [[1,3],[5,7],[9,10]]
    if (startTime > end) {
      count += startTime - end - (i === 0 ? 0 : 1);
    }
    if (endTime > end) {
      end = endTime;
    }
  }

  count += days - end;

  return count;
}
// @lc code=end
console.log(
  countDays(8, [
    [3, 4],
    [4, 8],
    [2, 5],
    [3, 8],
  ])
);

/*
// @lcpr case=start
// 8\n[[3,4],[4,8],[2,5],[3,8]]\n
// @lcpr case=end

// @lcpr case=start
// 5\n[[2,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// 6\n[[1,6]]\n
// @lcpr case=end

 */
