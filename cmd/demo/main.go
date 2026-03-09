package main

import (
	"fmt"
	"strings"
)

func startsOrEndsWithQuote(s string) bool {
	return strings.HasPrefix(s, "\"") || strings.HasPrefix(s, "'") ||
		strings.HasSuffix(s, "\"") || strings.HasSuffix(s, "'")
}

func main() {
	u := fmt.Sprintf("http://%s:%s/", "andrew", "8080")

	// 是否是使用单引号或者双引号的字符串
	if startsOrEndsWithQuote(u) {
		fmt.Println("URL must not begin or end with quotes")
	}

	fmt.Println(u)

}
