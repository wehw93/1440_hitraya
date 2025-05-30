package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	uniqueValues()
}

func uniqueValues() {
	arrUnique := make(map[string]bool)
	
	for n := 2; n < 100000; n++ {
		gn := g(n)
		key := fmt.Sprintf("%.12f", gn)
		arrUnique[key] = true
	}
	fmt.Printf("%d",len(arrUnique))
}
func g(n int) float64 {
	fn := f(n)
	ffn := f(fn)
	return float64(ffn) / float64(n)
}

func f(n int) int {
	str := strconv.Itoa(n)
	reversed := reverseString(str)
	reversed = strings.TrimLeft(reversed, "0")
	if reversed == "" {
		reversed = "0"
	}
	res, _ := strconv.Atoi(reversed)
	return res
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

