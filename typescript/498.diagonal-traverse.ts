/*
 * @lc app=leetcode.cn id=498 lang=typescript
 * @lcpr version=
 *
 * [498] 对角线遍历
 *
 * https://leetcode.cn/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (56.27%)
 * Likes:    532
 * Dislikes: 0
 * Total Accepted:    153.4K
 * Total Submissions: 267.6K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：mat = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,4,7,5,3,6,8,9]
 *
 *
 * 示例 2：
 *
 * 输入：mat = [[1,2],[3,4]]
 * 输出：[1,2,3,4]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == mat.length
 * n == mat[i].length
 * 1 <= m, n <= 10^4
 * 1 <= m * n <= 10^4
 * -10^5 <= mat[i][j] <= 10^5
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function findDiagonalOrder(mat: number[][]): number[] {
  const n = mat.length;
  const m = mat[0].length;

  const result: number[] = [];

  let x = 0,
    y = 0;
  while (true) {
    result.push(mat[x][y]);

    // 结束点判断
    if (x === n - 1 && y === m - 1) {
      break;
    }

    // 控制方向
    const dirt = (x + y) % 2 === 0 ? -1 : 1;
    let nextX = x + dirt;
    let nextY = y - dirt;

    // 边界判断
    if (nextX < 0 || nextX >= n || nextY < 0 || nextY >= m) {
      // 当前方向越界，需要转向
      if (dirt === -1) {
        // 右上方向越界
        if (y < m - 1) {
          // 优先向右移动
          y++;
        } else {
          // 向下移动
          x++;
        }
      } else {
        // 左下方向越界
        if (x < n - 1) {
          // 优先向下移动
          x++;
        } else {
          // 向右移动
          y++;
        }
      }
    } else {
      // 不越界，按原方向移动
      x = nextX;
      y = nextY;
    }
  }

  return result;
}
// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,4]]\n
// @lcpr case=end

 */
