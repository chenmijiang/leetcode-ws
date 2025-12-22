/*
 * @lc app=leetcode.cn id=54 lang=typescript
 * @lcpr version=30204
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (53.48%)
 * Likes:    2032
 * Dislikes: 0
 * Total Accepted:    884.9K
 * Total Submissions: 1.6M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2：
 *
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function spiralOrder(matrix: number[][]): number[] {
  const m = matrix.length;
  const n = matrix[0].length;
  const res: number[] = [];

  // 方向数组：右、下、左、上（顺时针）
  const dirs = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0],
  ];
  let dirIdx = 0;
  let row = 0,
    col = 0;

  for (let i = 0; i < m * n; i++) {
    res.push(matrix[row][col]);
    // 将访问过的元素标记为 101（超出范围 -100~100）
    matrix[row][col] = 101;

    // 计算下一个位置
    const nextRow = row + dirs[dirIdx][0];
    const nextCol = col + dirs[dirIdx][1];

    // 如果越界或已访问（值为101），则顺时针转向
    if (
      nextRow < 0 ||
      nextRow >= m ||
      nextCol < 0 ||
      nextCol >= n ||
      matrix[nextRow][nextCol] === 101
    ) {
      dirIdx = (dirIdx + 1) % 4;
    }

    // 移动到下一个位置
    row += dirs[dirIdx][0];
    col += dirs[dirIdx][1];
  }

  return res;
}
// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

 */
