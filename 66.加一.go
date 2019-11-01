import "strconv"

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (40.51%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    71.6K
 * Total Submissions: 176.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */
// @lc code=start
func plusOne(digits []int) []int {
  // 思路二
  index := len(digits) - 1
  carry := 1
  for i := index; i >= 0; i-- {
    num := digits[i]
    if num + carry <= 9 {
      digits[i] = num + carry
      return digits
    } else {
      if i == 0 {
        digits[i] = (num + carry) % 10
        carry = (num + carry) / 10
        temp := []int{carry}
        digits = append(temp, digits...)
        return digits
      }
      digits[i] = (num + carry) % 10
      carry = (num + carry) / 10
      continue
    }
  }
  return digits
}

// @lc code=end

// @Tips
/*
思路一
数组变成数字，相加之后，再变成数组,
思路错误，数组过长的时候，导致了不够存储
func plusOne(digits []int) []int {
	str := ""
	for i := 0; i < len(digits); i++ {
		str += strconv.Itoa(digits[i])
	}
	r := 0
	temp, err := strconv.Atoi(str)
	if err != nil {
		return nil
	} else {
		r = temp + 1
	}
	// 头部插入
	var result []int
  for r > 0 {   
    num := []int{r % 10}
		result = append(num, result...)
		r = r / 10
  }
	return result
}
思路二
只把最后一个变成数字，相加，考虑进位以及数组扩容
*/

