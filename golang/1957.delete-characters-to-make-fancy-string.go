/*
 * @lc app=leetcode.cn id=1957 lang=golang
 * @lcpr version=30204
 *
 * [1957] 删除字符使字符串变好
 *
 * https://leetcode.cn/problems/delete-characters-to-make-fancy-string/description/
 *
 * algorithms
 * Easy (68.54%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    29.2K
 * Total Submissions: 42.6K
 * Testcase Example:  '"leeetcode"'
 *
 * 一个字符串如果没有 三个连续 相同字符，那么它就是一个 好字符串 。
 *
 * 给你一个字符串 s ，请你从 s 删除 最少 的字符，使它变成一个 好字符串 。
 *
 * 请你返回删除后的字符串。题目数据保证答案总是 唯一的 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "leeetcode"
 * 输出："leetcode"
 * 解释：
 * 从第一组 'e' 里面删除一个 'e' ，得到 "leetcode" 。
 * 没有连续三个相同字符，所以返回 "leetcode" 。
 *
 *
 * 示例 2：
 *
 * 输入：s = "aaabaaaa"
 * 输出："aabaa"
 * 解释：
 * 从第一组 'a' 里面删除一个 'a' ，得到 "aabaaaa" 。
 * 从第二组 'a' 里面删除两个 'a' ，得到 "aabaa" 。
 * 没有连续三个相同字符，所以返回 "aabaa" 。
 *
 *
 * 示例 3：
 *
 * 输入：s = "aab"
 * 输出："aab"
 * 解释：没有连续三个相同字符，所以返回 "aab" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * s 只包含小写英文字母。
 *
 *
 */

// @lcpr-template-start
package main

import "strings"

// @lcpr-template-end
// @lc code=start
func makeFancyString(s string) string {
	n := len(s)
	if n <= 2 {
		return s
	}

	// 使用 strings.Builder 提高内存效率
	var builder strings.Builder
	builder.Grow(n)

	builder.WriteByte(s[0])
	count := 1

	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		// 只有当连续字符数不超过2时才添加
		if count <= 2 {
			builder.WriteByte(s[i])
		}
	}

	return builder.String()
}

// @lc code=end

/*
// @lcpr case=start
// "leeetcode"\n
// @lcpr case=end

// @lcpr case=start
// "aaabaaaa"\n
// @lcpr case=end

// @lcpr case=start
// "aab"\n
// @lcpr case=end

*/
