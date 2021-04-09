/*
 * @lc app=leetcode.cn id=451 lang=javascript
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var frequencySort = function (s) {
  // use map to set the count of each char in s
  let map = new Map();
  for (const c of s) {
    if (map.get(c)) {
      map.set(c, map.get(c) + 1);
    } else {
      map.set(c, 1);
    }
  }
  // map to array and sort
  let array = Array.from(map);
  array.sort((a, b) => b[1] - a[1]);
  let result = "";
  for (const v of array) {
    for (let i = 0; i < v[1]; i++) {
      result += v[0];
    }
  }
  return result;
};
// test case
console.log(frequencySort("tree"));
// @lc code=end
