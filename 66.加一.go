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
  var result []int
  for i := 0; i < len(strconv.Itoa(r)); i++ {
    result = result.append(r % 10, result...)
    r = r / 10
  }
}

