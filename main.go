package main

import (
  "fmt"
  // "reflect"
  // "strconv"
)

type ListNode struct {
  Val int
  Next *ListNode
}

// use this to debug local
func main() {
  var l1 *ListNode
  var l2 *ListNode
  l1 = &ListNode{2, &ListNode{4, &ListNode{4, nil}}}
  l2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
  fmt.Println(mergeTwoLists(l1, l2).Next.Next.Next.Val)
}

// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
  // 思路一，从头开始截取，然后进行判断
  if len(strs) > 0 {
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
  } else {
    return ""
  }
}
