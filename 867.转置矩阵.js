/*
 * @lc app=leetcode.cn id=867 lang=javascript
 *
 * [867] 转置矩阵
 */

// @lc code=start
/**
 * @param {number[][]} matrix
 * @return {number[][]}
 */
var transpose = function (matrix) {
  let result = []
  for (let i = 0; i < matrix.length; i++) {
    const array = matrix[i]
    for (let j = 0; j < array.length; j++) {
      const v = array[j]
      if (result[j]) {
        result[j].push(v)
      } else {
        result.push([v])
      }
    }
  }
  return result
}
// @lc code=end
