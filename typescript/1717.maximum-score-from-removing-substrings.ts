/*
 * @lc app=leetcode.cn id=1717 lang=typescript
 * @lcpr version=30204
 *
 * [1717] 删除子字符串的最大得分
 *
 * https://leetcode.cn/problems/maximum-score-from-removing-substrings/description/
 *
 * algorithms
 * Medium (50.20%)
 * Likes:    62
 * Dislikes: 0
 * Total Accepted:    14.2K
 * Total Submissions: 23.8K
 * Testcase Example:  '"cdbcbbaaabab"\n4\n5'
 *
 * 给你一个字符串 s 和两个整数 x 和 y 。你可以执行下面两种操作任意次。
 *
 *
 * 删除子字符串 "ab" 并得到 x 分。
 *
 *
 * 比方说，从 "cabxbae" 删除 ab ，得到 "cxbae" 。
 *
 *
 * 删除子字符串"ba" 并得到 y 分。
 *
 * 比方说，从 "cabxbae" 删除 ba ，得到 "cabxe" 。
 *
 *
 *
 *
 * 请返回对 s 字符串执行上面操作若干次能得到的最大得分。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "cdbcbbaaabab", x = 4, y = 5
 * 输出：19
 * 解释：
 * - 删除 "cdbcbbaaabab" 中加粗的 "ba" ，得到 s = "cdbcbbaaab" ，加 5 分。
 * - 删除 "cdbcbbaaab" 中加粗的 "ab" ，得到 s = "cdbcbbaa" ，加 4 分。
 * - 删除 "cdbcbbaa" 中加粗的 "ba" ，得到 s = "cdbcba" ，加 5 分。
 * - 删除 "cdbcba" 中加粗的 "ba" ，得到 s = "cdbc" ，加 5 分。
 * 总得分为 5 + 4 + 5 + 5 = 19 。
 *
 * 示例 2：
 *
 * 输入：s = "aabbaaxybbaabb", x = 5, y = 4
 * 输出：20
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^5
 * 1 <= x, y <= 10^4
 * s 只包含小写英文字母。
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function maximumGain(s: string, x: number, y: number): number {
  // 贪心优先处理得分更高的子串
  const [highTarget, highScore, lowTarget, lowScore] =
    x >= y ? ['ab', x, 'ba', y] : ['ba', y, 'ab', x];

  const stack: string[] = [];
  let totalScore = 0;

  // 第一轮处理：优先消除高分子串
  for (const c of s) {
    stack.push(c);
    if (stack.length >= 2 && stack[stack.length - 2] + stack[stack.length - 1] === highTarget) {
      stack.pop();
      stack.pop();
      totalScore += highScore;
    }
  }

  // 第二轮处理：处理剩下的 lowTarget
  const newStack: string[] = [];
  for (const c of stack) {
    newStack.push(c);
    if (
      newStack.length >= 2 &&
      newStack[newStack.length - 2] + newStack[newStack.length - 1] === lowTarget
    ) {
      newStack.pop();
      newStack.pop();
      totalScore += lowScore;
    }
  }

  return totalScore;
}

// @lc code=end

console.log(maximumGain('cdbcbbaaabab', 4, 5) === 19);
console.log(maximumGain('aabbaaxybbaabb', 5, 4) === 20);

/*
// @lcpr case=start
// "cdbcbbaaabab"\n4\n5\n
// @lcpr case=end

// @lcpr case=start
// "aabbaaxybbaabb"\n5\n4\n
// @lcpr case=end

 */
