/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (47.70%)
 * Likes:    1297
 * Dislikes: 0
 * Total Accepted:    110.2K
 * Total Submissions: 230.7K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 * 
 * 示例:
 * 
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 * 
 * 
 * 进阶:
 * 
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 * 
 */

// @lc code=start
func maxSubArray(nums []int) int {
	length := len(nums)
	result := nums[0]
  for i := length; i > 0; i-- {
		for j := 0; j < length - i + 1; j++ {
			temps := nums[j:j+i]
			tempResult := 0
			for _, temp := range temps {
				tempResult += temp
			}
			if tempResult > result {
				result = tempResult
			}
		}
	}
	return result
}
// @lc code=end

// @Tips
/*
时间限制问题，时间复杂度超了。
*/