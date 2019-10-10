package main

import "fmt"

// use this to debug local
func main() {
  strs := []string{"flower","flow",""}
  fmt.Println(longestCommonPrefix(strs))
}

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
