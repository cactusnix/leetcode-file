import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (32.88%)
 * Likes:    1257
 * Dislikes: 0
 * Total Accepted:    170.8K
 * Total Submissions: 519.2K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}
	var result int64 = 0
	str := strconv.Itoa(x)
	for i := len(str) - 1; i >= 0; i-- {
		temp := int64(x % 10)
		for j := 0; j < i; j++ {
			temp = temp * 10
		}
		result += temp
		x = x / 10
	}
	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}
	return int(result) * sign
}

