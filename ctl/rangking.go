package main

import (
	"fmt"
	"strconv"
	"strings"
)

// getRanking makes this method more elegant
func getRanking() int {
	// Get web page data
	URL := fmt.Sprintf("https://leetcode.com/%s/", getConfig().Username)
	data := getRaw(URL)
	str := string(data)
	// Obtain ranking information by continuously cutting str
	fmt.Println(str)
	i := strings.Index(str, "ng-init")
	j := i + strings.Index(str[i:], "ng-cloak")
	str = str[i:j]
	i = strings.Index(str, "(")
	j = strings.Index(str, ")")
	str = str[i:j]
	//	fmt.Println("2\n", str)
	strs := strings.Split(str, ",")
	str = strs[6]
	//	fmt.Println("1\n", str)
	i = strings.Index(str, "'")
	j = 2 + strings.Index(str[2:], "'")
	//	fmt.Println("0\n", str)
	str = str[i+1 : j]
	r, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Unable to convert %s to numeric Ranking", str)
	}
	return r
}
