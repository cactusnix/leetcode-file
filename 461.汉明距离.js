/*
 * @lc app=leetcode.cn id=461 lang=javascript
 *
 * [461] 汉明距离
 */

// @lc code=start
/**
 * @param {number} x
 * @param {number} y
 * @return {number}
 */

/**
 * 汉明距离广泛用于多个领域。
 * 在编码理论中用于错误检测，在信息论中量化自负床之间的差异。
 * 两个整数之间的汉明距离是对应位置上数字不同的位数。
 */
var hammingDistance = function (x, y) {
  let s = x ^ y,
    ret = 0;
  while (s != 0) {
    ret += s & 1;
    s >>= 1;
  }
  return ret;
};

console.log(hammingDistance(1, 4));
// @lc code=end
