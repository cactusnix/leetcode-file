/*
 * @lc app=leetcode.cn id=283 lang=javascript
 *
 * [283] 移动零
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */

// 憨憨方法
// var moveZeroes = function(nums) {
//   let length = nums.length
//   for (let i = 0; i < length; i++) {
//     if (nums[i] === 0) {
//       nums.push(nums[i])
//       nums.splice(i, 1)
//       i--
//       length--
//     }
//   }
// };

// 双指针
var moveZeroes = function(nums) {
  let left = right = 0
  let len = nums.length
  while (right < len) {
    if (nums[right] != 0) {
      const temp = nums[right]
      nums[right] = nums[left]
      nums[left] = temp
      left++
    }
    right++
  }
};
// @lc code=end

