package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	isSecured := securePassword("asd123asD")
	fmt.Println(isSecured)
}

func securePassword(pass string) bool {
	return utf8.RuneCountInString(pass) >= 6 &&
		pass != strings.ToLower(pass) &&
		pass != strings.ToUpper(pass)

}
