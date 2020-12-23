/*
 * @lc app=leetcode.cn id=387 lang=javascript
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
var firstUniqChar = function (s) {
  for (let i = 0; i < s.length; i++) {
    const v = s[i]
    if (s.slice(0, i).indexOf(v) === -1 && s.slice(i + 1).indexOf(v) === -1) {
      return i
    }
  }
  return -1
}
// @lc code=end
