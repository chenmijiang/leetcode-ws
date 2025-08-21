/*
 * @lc app=leetcode.cn id=1504 lang=typescript
 * @lcpr version=30204
 *
 * [1504] 统计全 1 子矩形
 *
 * https://leetcode.cn/problems/count-submatrices-with-all-ones/description/
 *
 * algorithms
 * Medium (64.95%)
 * Likes:    255
 * Dislikes: 0
 * Total Accepted:    25.3K
 * Total Submissions: 36.8K
 * Testcase Example:  '[[1,0,1],[1,1,0],[1,1,0]]'
 *
 * 给你一个 m x n 的二进制矩阵 mat ，请你返回有多少个 子矩形 的元素全部都是 1 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：mat = [[1,0,1],[1,1,0],[1,1,0]]
 * 输出：13
 * 解释：
 * 有 6 个 1x1 的矩形。
 * 有 2 个 1x2 的矩形。
 * 有 3 个 2x1 的矩形。
 * 有 1 个 2x2 的矩形。
 * 有 1 个 3x1 的矩形。
 * 矩形数目总共 = 6 + 2 + 3 + 1 + 1 = 13 。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]
 * 输出：24
 * 解释：
 * 有 8 个 1x1 的子矩形。
 * 有 5 个 1x2 的子矩形。
 * 有 2 个 1x3 的子矩形。
 * 有 4 个 2x1 的子矩形。
 * 有 2 个 2x2 的子矩形。
 * 有 2 个 3x1 的子矩形。
 * 有 1 个 3x2 的子矩形。
 * 矩形数目总共 = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24 。
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 150
 * mat[i][j] 仅包含 0 或 1
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function numSubmat(mat: number[][]): number {
  const m = mat.length;
  const n = mat[0].length;

  let ans = 0;
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (mat[i][j] === 1) {
        mat[i][j] = (i > 0 ? mat[i - 1][j] : 0) + 1;
        let h = mat[i][j];
        for (let k = j; k >= 0 && mat[i][k] > 0; k--) {
          h = Math.min(h, mat[i][k]);
          ans += h;
        }
      }
    }
  }
  return ans;
}
// @lc code=end

/*
// @lcpr case=start
// [[1,0,1],[1,1,0],[1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,1,0],[0,1,1,1],[1,1,1,0]]\n
// @lcpr case=end

 */
