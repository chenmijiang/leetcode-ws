/*
 * @lc app=leetcode.cn id=2410 lang=typescript
 * @lcpr version=30204
 *
 * [2410] 运动员和训练师的最大匹配数
 *
 * https://leetcode.cn/problems/maximum-matching-of-players-with-trainers/description/
 *
 * algorithms
 * Medium (68.53%)
 * Likes:    38
 * Dislikes: 0
 * Total Accepted:    20K
 * Total Submissions: 27.4K
 * Testcase Example:  '[4,7,9]\n[8,2,5,8]'
 *
 * 给你一个下标从 0 开始的整数数组 players ，其中 players[i] 表示第 i 名运动员的 能力 值，同时给你一个下标从 0
 * 开始的整数数组 trainers ，其中 trainers[j] 表示第 j 名训练师的 训练能力值 。
 *
 * 如果第 i 名运动员的能力值 小于等于 第 j 名训练师的能力值，那么第 i 名运动员可以 匹配 第 j
 * 名训练师。除此以外，每名运动员至多可以匹配一位训练师，每位训练师最多可以匹配一位运动员。
 *
 * 请你返回满足上述要求 players 和 trainers 的 最大 匹配数。
 *
 *
 *
 * 示例 1：
 *
 * 输入：players = [4,7,9], trainers = [8,2,5,8]
 * 输出：2
 * 解释：
 * 得到两个匹配的一种方案是：
 * - players[0] 与 trainers[0] 匹配，因为 4 <= 8 。
 * - players[1] 与 trainers[3] 匹配，因为 7 <= 8 。
 * 可以证明 2 是可以形成的最大匹配数。
 *
 *
 * 示例 2：
 *
 * 输入：players = [1,1,1], trainers = [10]
 * 输出：1
 * 解释：
 * 训练师可以匹配所有 3 个运动员
 * 每个运动员至多只能匹配一个训练师，所以最大答案是 1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= players.length, trainers.length <= 10^5
 * 1 <= players[i], trainers[j] <= 10^9
 *
 *
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
function matchPlayersAndTrainers(players: number[], trainers: number[]): number {
  players.sort((a, b) => a - b);
  trainers.sort((a, b) => a - b);

  let playerIndex = 0;
  let trainerIndex = 0;
  let matchCount = 0;

  if (players[0] > trainers[trainers.length - 1]) {
    return 0;
  }

  while (playerIndex < players.length && trainerIndex < trainers.length) {
    if (players[playerIndex] <= trainers[trainerIndex]) {
      matchCount++;
      playerIndex++;
    }
    trainerIndex++;
  }

  return matchCount;
}
// @lc code=end

/*
// @lcpr case=start
// [4,7,9]\n[8,2,5,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n[10]\n
// @lcpr case=end

// @lcpr case=start
// [56563725,763661320,707284353,169559687,824944854]\n[764327984,472832686,732262503,412251220,567505886,514120975,592748479,377842496,705222514,357754473,226699700,376370669,700431674,549362887,572373629,29907462,455099822,697035225,752841629,868901334,456744149,390681662,99669290,91713465,118141670]\n
// @lcpr case=end

 */
