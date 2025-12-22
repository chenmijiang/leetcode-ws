/*
 * @lc app=leetcode.cn id=54 lang=rust
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
use crate::Solution;
// @lcpr-template-end
// @lc code=start
impl Solution {
    pub fn spiral_order(mut matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let m = matrix.len();
        let n = matrix[0].len();
        let mut res = Vec::with_capacity(m * n);

        let dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)];
        let mut dir_idx = 0;
        let mut row = 0i32;
        let mut col = 0i32;

        for _ in 0..m * n {
            res.push(matrix[row as usize][col as usize]);
            // 将访问过的元素标记为 101（超出范围 -100~100）
            matrix[row as usize][col as usize] = 101;

            let next_row = row + dirs[dir_idx].0;
            let next_col = col + dirs[dir_idx].1;

            // 如果越界或已访问（值为101），则顺时针转向
            if next_row < 0
                || next_row >= m as i32
                || next_col < 0
                || next_col >= n as i32
                || matrix[next_row as usize][next_col as usize] == 101
            {
                dir_idx = (dir_idx + 1) % 4;
            }

            row += dirs[dir_idx].0;
            col += dirs[dir_idx].1;
        }

        res
    }
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
