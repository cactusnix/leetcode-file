/*
 * @lc app=leetcode.cn id=993 lang=javascript
 *
 * [993] 二叉树的堂兄弟节点
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
 * @param {number} x
 * @param {number} y
 * @return {boolean}
 */
var isCousins = function (root, x, y) {
  let x_parent = null,
    x_depth = null,
    x_found = false;
  let y_parent = null,
    y_depth = null,
    y_found = false;
  let update = (node, parent, depth) => {
    if (node.val === x) {
      [x_parent, x_depth, x_found] = [parent, depth, true];
    }
    if (node.val === y) {
      [y_parent, y_depth, y_found] = [parent, depth, true];
    }
  };
  let queue = [[root, 0]];
  update(root, null, 0);
  while (queue.length) {
    const [node, depth] = queue.shift();
    if (node.left) {
      queue.push([node.left, depth + 1]);
      update(node.left, node, depth + 1);
    }
    if (node.right) {
      queue.push([node.right, depth + 1]);
      update(node.right, node, depth + 1);
    }
    if (x_found && y_found) {
      break;
    }
  }
  // parent 的不等于是对象的不等于
  return x_depth === y_depth && x_parent !== y_parent;
};
// @lc code=end
