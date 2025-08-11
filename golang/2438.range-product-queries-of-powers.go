/*
 * @lc app=leetcode.cn id=2438 lang=golang
 * @lcpr version=30204
 *
 * [2438] 二的幂数组中查询范围内的乘积
 *
 * https://leetcode.cn/problems/range-product-queries-of-powers/description/
 *
 * algorithms
 * Medium (44.51%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    18.2K
 * Total Submissions: 35.1K
 * Testcase Example:  '15\n[[0,1],[2,2],[0,3]]'
 *
 * 给你一个正整数 n ，你需要找到一个下标从 0 开始的数组 powers ，它包含 最少 数目的 2 的幂，且它们的和为 n 。powers 数组是
 * 非递减 顺序的。根据前面描述，构造 powers 数组的方法是唯一的。
 *
 * 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] ，其中
 * queries[i] 表示请你求出满足 lefti <= j <= righti 的所有 powers[j] 的乘积。
 *
 * 请你返回一个数组 answers ，长度与 queries 的长度相同，其中 answers[i]是第 i
 * 个查询的答案。由于查询的结果可能非常大，请你将每个 answers[i] 都对 10^9 + 7 取余 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 15, queries = [[0,1],[2,2],[0,3]]
 * 输出：[2,4,64]
 * 解释：
 * 对于 n = 15 ，得到 powers = [1,2,4,8] 。没法得到元素数目更少的数组。
 * 第 1 个查询的答案：powers[0] * powers[1] = 1 * 2 = 2 。
 * 第 2 个查询的答案：powers[2] = 4 。
 * 第 3 个查询的答案：powers[0] * powers[1] * powers[2] * powers[3] = 1 * 2 * 4 * 8 =
 * 64 。
 * 每个答案对 10^9 + 7 取余得到的结果都相同，所以返回 [2,4,64] 。
 *
 *
 * 示例 2：
 *
 * 输入：n = 2, queries = [[0,0]]
 * 输出：[2]
 * 解释：
 * 对于 n = 2, powers = [2] 。
 * 唯一一个查询的答案是 powers[0] = 2 。答案对 10^9 + 7 取余后结果相同，所以返回 [2] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^9
 * 1 <= queries.length <= 10^5
 * 0 <= starti <= endi < powers.length
 *
 *
 */

// @lcpr-template-start
package main

// @lcpr-template-end
// @lc code=start
func productQueries(n int, queries [][]int) []int {
	MOD := 1000000007
	powers := []int{}
	power := 1
	for power <= n {
		if n&power > 0 {
			powers = append(powers, power)
		}
		power <<= 1
	}

	// 构建前缀积数组
	prefix := make([]int, len(powers)+1)
	prefix[0] = 1
	for i := 1; i <= len(powers); i++ {
		prefix[i] = (prefix[i-1] * powers[i-1]) % MOD
	}

	// 快速幂算法计算模逆元
	var pow func(base, exp, mod int) int
	pow = func(base, exp, mod int) int {
		result := 1
		base %= mod
		for exp > 0 {
			if exp&1 == 1 {
				result = (result * base) % mod
			}
			base = (base * base) % mod
			exp >>= 1
		}
		return result
	}

	answers := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		// 使用前缀积计算区间乘积: prefix[right+1] / prefix[left]
		// 在模运算下，除法转换为乘以模逆元
		inv := pow(prefix[left], MOD-2, MOD)
		answers[i] = (prefix[right+1] * inv) % MOD
	}
	return answers
}

// @lc code=end

/*
// @lcpr case=start
// 15\n[[0,1],[2,2],[0,3]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[0,0]]\n
// @lcpr case=end

*/
