/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (31.29%)
 * Likes:    137
 * Dislikes: 0
 * Total Accepted:    48.9K
 * Total Submissions: 155.8K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	sArray := strings.Split(s, " ")
	return len(sArray[len(sArray)-1])
}

// @lc code=end

// @Tips
/*
思路一
去除尾部空格，然后按空格分割成数组，最后一个数组的长度
思路二
不转换，直接获取最后一个单词
 */
