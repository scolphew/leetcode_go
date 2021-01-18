package main

import "fmt"

func letterCasePermutation(S string) []string {
	ans := make([]string, 0)
	ans = append(ans, S)
	for i, c := range S {
		l := len(ans)
		if c >= 'a' && c <= 'z' {
			for j := 0; j < l; j++ {
				tmp := []byte(ans[j])
				tmp[i] = ans[j][i] - 32
				ans = append(ans, string(tmp))
			}
		}
		if c >= 'A' && c <= 'Z' {
			for j := 0; j < l; j++ {
				tmp := []byte(ans[j])
				tmp[i] = ans[j][i] + 32
				ans = append(ans, string(tmp))
			}
		}
	}
	return ans
}

func main() {
	strings := letterCasePermutation("abcde")
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
