/*
 * @lc app=leetcode.cn id=1046 lang=javascript
 *
 * [1046] 最后一块石头的重量
 */

// @lc code=start
/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function (stones) {
  stones.sort((a, b) => a - b)
  for (let i = stones.length - 1; i >= 0; i--) {
    if (i != 0) {
      if (stones[i] === stones[i - 1]) {
        stones.pop()
        stones.pop()
        i = i - 1
      } else {
        stones[i - 1] = stones[i] - stones[i - 1]
        stones.pop()
        stones.sort((a, b) => a - b)
      }
    }
  }
  if (stones.length > 0) {
    return stones[0]
  } else {
    return 0
  }
}
// @lc code=end
