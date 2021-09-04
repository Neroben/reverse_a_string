package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := input()
	fmt.Printf("string before = %s, string after = %s", str, reverseString(str))
}

func input() string {
	fmt.Printf("Введите число для проверки: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func reverseString(str string) string {
	strArr := []rune(str)
	for i, j := 0, len(str)-1; i < j; i,j = i + 1, j - 1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)
}