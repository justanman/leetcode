package main

import (
  "fmt"
  "strings"
)

func longestCommonPrefix(strs []string) string {
  lcp := ""
  if len(strs) == 0 {
    return lcp
  }
  lcp = strs[0]
  for i := 0; i < len(strs); i++ {
    for strings.Index(strs[i], lcp) != 0 {
      lcp = lcp[0:len(lcp) - 1]
      if len(lcp) == 0 {
        return ""
      }
    }
  }
  return lcp
}

func main() {
  fmt.Println(longestCommonPrefix([]string{"abc", "abd", "abcde"}) == "ab")
}
