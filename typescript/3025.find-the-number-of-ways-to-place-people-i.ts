/*
 * @lc app=leetcode.cn id=3025 lang=typescript
 * @lcpr version=
 *
 * [3025] 人员站位的方案数 I
 *
 * https://leetcode.cn/problems/find-the-number-of-ways-to-place-people-i/description/
 *
 * algorithms
 * Medium (46.66%)
 * Likes:    23
 * Dislikes: 0
 * Total Accepted:    11.2K
 * Total Submissions: 19.6K
 * Testcase Example:  '[[1,1],[2,2],[3,3]]'
 *
 * 给你一个  n x 2 的二维数组 points ，它表示二维平面上的一些点坐标，其中 points[i] = [xi, yi] 。
 *
 *
 *
 * 计算点对 (A, B) 的数量，其中
 *
 *
 * A 在 B 的左上角，并且
 * 它们形成的长方形中（或直线上）没有其它点（包括边界）。
 *
 *
 * 返回数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：points = [[1,1],[2,2],[3,3]]
 *
 * 输出：0
 *
 * 解释：
 *
 *
 *
 * 没有办法选择 A 和 B，使得 A 在 B 的左上角。
 *
 *
 * 示例 2：
 *
 *
 * 输入：points = [[6,2],[4,4],[2,6]]
 *
 * 输出：2
 *
 * 解释：
 *
 *
 *
 *
 * 左边的是点对 (points[1], points[0])，其中 points[1] 在 points[0]
 * 的左上角，并且形成的长方形内部是空的。
 * 中间的是点对 (points[2], points[1])，和左边的一样是合法的点对。
 * 右边的是点对 (points[2], points[0])，其中 points[2] 在 points[0] 的左上角，但 points[1]
 * 在长方形内部，所以不是一个合法的点对。
 *
 *
 *
 * 示例 3：
 *
 *
 * 输入：points = [[3,1],[1,3],[1,1]]
 *
 * 输出：2
 *
 * 解释：
 *
 *
 *
 *
 * 左边的是点对 (points[2], points[0])，其中 points[2] 在 points[0]
 * 的左上角并且在它们形成的直线上没有其它点。注意两个点形成一条线的情况是合法的。
 * 中间的是点对 (points[1], points[2])，和左边一样也是合法的点对。
 * 右边的是点对 (points[1], points[0])，它不是合法的点对，因为 points[2] 在长方形的边上。
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 50
 * points[i].length == 2
 * 0 <= points[i][0], points[i][1] <= 50
 * points[i] 点对两两不同。
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function numberOfPairs(points: number[][]): number {
  const n = points.length;
  let total = 0;

  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (i === j) continue;

      const A = points[i];
      const B = points[j];

      // 检查 A 是否在 B 的左上角
      // A 在 B 的左上角意味着：A.x <= B.x 且 A.y >= B.y
      if (A[0] > B[0] || A[1] < B[1]) {
        continue;
      }

      // 检查矩形内部是否有其他点
      let isValid = true;
      // 矩形边界：[min(A.x, B.x), max(A.x, B.x)] × [min(A.y, B.y), max(A.y, B.y)]
      const minX = Math.min(A[0], B[0]);
      const maxX = Math.max(A[0], B[0]);
      const minY = Math.min(A[1], B[1]);
      const maxY = Math.max(A[1], B[1]);

      for (let k = 0; k < n; k++) {
        if (k === i || k === j) continue;

        // 检查点 C 是否在矩形内部
        const C = points[k];

        // 如果点 C 在矩形内部（包括边界），则这个点对无效
        if (C[0] >= minX && C[0] <= maxX && C[1] >= minY && C[1] <= maxY) {
          isValid = false;
          break;
        }
      }

      if (isValid) {
        total++;
      }
    }
  }

  return total;
}
// @lc code=end

console.log(
  numberOfPairs([
    [0, 1],
    [0, 2],
    [0, 4],
  ])
);

/*
// @lcpr case=start
// [[1,1],[2,2],[3,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[6,2],[4,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,1],[1,3],[1,1]]\n
// @lcpr case=end

 */
