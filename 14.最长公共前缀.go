/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.53%)
 * Likes:    720
 * Dislikes: 0
 * Total Accepted:    131.9K
 * Total Submissions: 377.6K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
   // 思路一，从头开始截取，然后进行判断
   if len(strs) > 0 {
    if len(strs) == 1 {
      return strs[0]
    } else {
      for _, str := range strs {
        if str == "" {
          return ""
        }
      }
      if strs[0][:1] != strs[1][:1] {
        return ""
      } else {
        length := 1
        // 求得最短长度
        minLength := len(strs[0])
        for i := 0; i < len(strs); i++ {
          if minLength < len(strs[i]) {
            continue
          } else {
            minLength = len(strs[i])
          }
        }
        prefix := strs[0][:length]
        for i := 0; i < minLength; i++ {
          prefix = strs[0][:length]
          for j := 0; j < len(strs); j++ {
            if prefix != strs[j][:length] {
              return prefix[:len(prefix)-1]
            } else {
              continue
            }
          }
          length = length+1
        }
        return prefix
      }
    }
  } else {
    return ""
  }
}
// @lc code=end

// @Tips
/*
虽然是到easy题目，但是却要考虑很多的testcase
比如空数组，空字符串
比如数组长度为一
这些考虑基于我写的代码思路一
*/
