/*
 * @lc app=leetcode.cn id=1006 lang=javascript
 *
 * [1006] 笨阶乘
 */

// @lc code=start
/**
 * @param {number} N
 * @return {number}
 */
var clumsy = function (N) {
  let stack = [N];
  N--;
  let i = 0;
  while (N > 0) {
    if (i % 4 === 0) {
      stack.push(stack.pop() * N);
    } else if (i % 4 === 1) {
      const value = stack.pop() / N;
      stack.push(value > 0 ? Math.floor(value) : Math.ceil(value));
    } else if (i % 4 === 2) {
      stack.push(N);
    } else {
      stack.push(-N);
    }
    i++;
    N--;
  }
  let sum = 0;
  stack.forEach((v) => {
    sum += v;
  });
  return sum;
};
console.log(clumsy(10));
// @lc code=end
