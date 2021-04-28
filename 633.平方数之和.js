/*
 * @lc app=leetcode.cn id=633 lang=javascript
 *
 * [633] 平方数之和
 */

// @lc code=start
/**
 * @param {number} c
 * @return {boolean}
 */
// 暴力枚举
/* var judgeSquareSum = function (c) {
  // a * a <= c 可以这么优化
  for (let a = 0; a <= Math.floor(Math.sqrt(c)); a++) {
    const b = Math.sqrt(c - a * a);
    if (b === parseInt(b)) {
      return true;
    }
  }
  return false;
};
*/
// 双指针
var judgeSquareSum = function (c) {
  let left = 0,
    right = Math.floor(Math.sqrt(c));
  while (left <= right) {
    const temp = left * left + right * right;
    if (temp === c) {
      return true;
    }
    if (temp < c) {
      left++;
    }
    if (temp > c) {
      right--;
    }
  }
  return false;
};
console.log(judgeSquareSum(2));
// @lc code=end
