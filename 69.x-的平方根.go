/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.06%)
 * Likes:    254
 * Dislikes: 0
 * Total Accepted:    72.4K
 * Total Submissions: 194.4K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 * 
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 * 
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 * 
 * 示例 1:
 * 
 * 输入: 4
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842..., 
 * 由于返回类型是整数，小数部分将被舍去。
 * 
 * 
 */

// @lc code=start
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 1
	right := x / 2 + 1
	for left < right {
		mid := (left + right)	/ 2 + 1
		if mid * mid > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
// @lc code=end

// @Tips
/*
* 二分查找
*/
