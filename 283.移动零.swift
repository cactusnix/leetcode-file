/*
 * @lc app=leetcode.cn id=283 lang=swift
 *
 * [283] 移动零
 */

// @lc code=start
class Solution {
    func moveZeroes(_ nums: inout [Int]) {
        var left = 0, right = 0
        var len = nums.count
        while right < len {
            if nums[right] != 0 {
                nums.swapAt(left, right)
                left = left + 1
            }
            right = right + 1
        }
    }
}
// @lc code=end

