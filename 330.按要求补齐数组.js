/*
 * @lc app=leetcode.cn id=330 lang=javascript
 *
 * [330] 按要求补齐数组
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} n
 * @return {number}
 */
var minPatches = function (nums, n) {
  let i = 0,
    m = 1,
    r = 0
  while (m <= n) nums[i] <= m ? (m += nums[i++]) : ((m += m), r++)
  return r
}
// @lc code=end
