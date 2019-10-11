/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (56.72%)
 * Likes:    656
 * Dislikes: 0
 * Total Accepted:    121.5K
 * Total Submissions: 210.9K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  if l1 == nil && l2 != nil {
    return l2
  }
  if l1 != nil && l2 == nil {
    return l1
  }
  if l1 == nil && l2 == nil {
    return nil
  }
  var l3 *ListNode
  var l4 *ListNode
  if l2.Val >= l1.Val {
    l3 = l1
    l4 = l2
  } else {
    l3 = l2
    l4 = l1
  }
  for l4 != nil {
    fmt.Println("l2 Val", l4.Val)
    // 不改变l1的指向，这样就可以返回，这里是引用，所以会修改l1的指向
    l5 := l3
    for l5 != nil {
      fmt.Println("l1 Val", l5.Val)
      if l4.Val >= l5.Val {
        if l5.Next == nil {
          fmt.Println("当前的l1", l5.Val)
          temp := &ListNode{l4.Val, nil}
          l5.Next = temp
          break
        } else if l4.Val <= l5.Next.Val {
          fmt.Println("当前的l1", l5.Val)
          temp := &ListNode{l4.Val, l5.Next}
          l5.Next = temp
          break
        }
      }
      l5 = l5.Next
    }
    l4 = l4.Next
  }
  return l3
}
// @lc code=end

// @Tips
/*
思路是正确的，却被指针的坑搞晕了。
指针真的搞晕了。内存占用过大，不停拷贝指针的副本。
暂存
*/
