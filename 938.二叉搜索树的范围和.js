/*
 * @lc app=leetcode.cn id=938 lang=javascript
 *
 * [938] 二叉搜索树的范围和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var rangeSumBST = function (root, low, high) {
  let result = 0;
  let preorder = (root) => {
    if (!root) {
      return;
    }
    if (root.val >= low && root.val <= high) {
      result += root.val;
    }
    preorder(root.left);
    preorder(root.right);
  };
  preorder(root);
  return result;
};
console.log(
  rangeSumBST(
    {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
        left: {
          val: 4,
          left: {
            val: 6,
          },
          right: {
            val: 7,
          },
        },
      },
    },
    2,
    5
  )
);
// @lc code=end
