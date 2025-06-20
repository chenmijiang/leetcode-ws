/*
 * @lc app=leetcode.cn id=3443 lang=golang
 * @lcpr version=30204
 *
 * [3443] K 次修改后的最大曼哈顿距离
 *
 * https://leetcode.cn/problems/maximum-manhattan-distance-after-k-changes/description/
 *
 * algorithms
 * Medium (47.77%)
 * Likes:    25
 * Dislikes: 0
 * Total Accepted:    10.7K
 * Total Submissions: 17.5K
 * Testcase Example:  '"NWSE"\n1'
 *
 * 给你一个由字符 'N'、'S'、'E' 和 'W' 组成的字符串 s，其中 s[i] 表示在无限网格中的移动操作：
 *
 *
 * 'N'：向北移动 1 个单位。
 * 'S'：向南移动 1 个单位。
 * 'E'：向东移动 1 个单位。
 * 'W'：向西移动 1 个单位。
 *
 *
 * 初始时，你位于原点 (0, 0)。你 最多 可以修改 k 个字符为任意四个方向之一。
 *
 * 请找出在 按顺序 执行所有移动操作过程中的 任意时刻 ，所能达到的离原点的 最大曼哈顿距离 。
 *
 * 曼哈顿距离 定义为两个坐标点 (xi, yi) 和 (xj, yj) 的横向距离绝对值与纵向距离绝对值之和，即 |xi - xj| + |yi -
 * yj|。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "NWSE", k = 1
 *
 * 输出：3
 *
 * 解释：
 *
 * 将 s[2] 从 'S' 改为 'N' ，字符串 s 变为 "NWNE" 。
 *
 *
 *
 *
 * 移动操作
 * 位置 (x, y)
 * 曼哈顿距离
 * 最大值
 *
 *
 *
 *
 * s[0] == 'N'
 * (0, 1)
 * 0 + 1 = 1
 * 1
 *
 *
 * s[1] == 'W'
 * (-1, 1)
 * 1 + 1 = 2
 * 2
 *
 *
 * s[2] == 'N'
 * (-1, 2)
 * 1 + 2 = 3
 * 3
 *
 *
 * s[3] == 'E'
 * (0, 2)
 * 0 + 2 = 2
 * 3
 *
 *
 *
 *
 * 执行移动操作过程中，距离原点的最大曼哈顿距离是 3 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "NSWWEW", k = 3
 *
 * 输出：6
 *
 * 解释：
 *
 * 将 s[1] 从 'S' 改为 'N' ，将 s[4] 从 'E' 改为 'W' 。字符串 s 变为 "NNWWWW" 。
 *
 * 执行移动操作过程中，距离原点的最大曼哈顿距离是 6 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * 0 <= k <= s.length
 * s 仅由 'N'、'S'、'E' 和 'W' 。
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start

func maxDistance(s string, k int) int {
	x, y := 0, 0 // 当前坐标
	maxDist := 0

	// 只需要一次遍历，直接计算坐标
	for i, ch := range s {
		switch ch {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}

		// 当前曼哈顿距离
		currDist := abs(x) + abs(y)

		// 在当前步数下的理论最大距离
		theoreticalMax := i + 1

		// k次修改能达到的最大距离
		possibleDist := currDist + 2*k

		// 取较小值，因为不能超过步数
		maxDist = max(maxDist, min(possibleDist, theoreticalMax))
	}

	return maxDist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "NWSE"\n1\n
// @lcpr case=end

// @lcpr case=start
// "NSWWEW"\n3\n
// @lcpr case=end

*/
