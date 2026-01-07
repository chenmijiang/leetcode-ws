/*
 * @lc app=leetcode.cn id=1339 lang=typescript
 * @lcpr version=30204
 *
 * [1339] 分裂二叉树的最大乘积
 *
 * https://leetcode.cn/problems/maximum-product-of-splitted-binary-tree/description/
 *
 * algorithms
 * Medium (43.11%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    27.4K
 * Total Submissions: 56.9K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。
 *
 * 由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：root = [1,2,3,4,5,6]
 * 输出：110
 * 解释：删除红色的边，得到 2 棵子树，和分别为 11 和 10 。它们的乘积是 110 （11*10）
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：root = [1,null,2,3,4,null,null,5,6]
 * 输出：90
 * 解释：移除红色的边，得到 2 棵子树，和分别是 15 和 6 。它们的乘积为 90 （15*6）
 *
 *
 * 示例 3：
 *
 * 输入：root = [2,3,9,10,7,8,6,5,4,11,1]
 * 输出：1025
 *
 *
 * 示例 4：
 *
 * 输入：root = [1,1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每棵树最多有 50000 个节点，且至少有 2 个节点。
 * 每个节点的值在 [1, 10000] 之间。
 *
 *
 */

// @lcpr-template-start
class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}
// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function maxProduct(root: TreeNode | null): number {
  const MOD = 1e9 + 7;
  let total = 0;
  let maxProd = 0;

  function getSum(node: TreeNode | null): number {
    if (!node) return 0;
    return node.val + getSum(node.left) + getSum(node.right);
  }

  total = getSum(root);

  function findMaxProd(node: TreeNode | null): number {
    if (!node) return 0;

    const subSum = node.val + findMaxProd(node.left) + findMaxProd(node.right);
    maxProd = Math.max(maxProd, subSum * (total - subSum));
    return subSum;
  }

  findMaxProd(root);
  return maxProd % MOD;
}
// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2,3,4,null,null,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,9,10,7,8,6,5,4,11,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

 */
