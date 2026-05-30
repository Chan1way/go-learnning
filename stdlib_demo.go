package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stdlibDemo() {
	// strings 字符串处理
	fmt.Println(strings.ToUpper("hello"))         // HELLO
	fmt.Println(strings.Contains("Hello", "ell")) // true
	fmt.Println(strings.Split("a,b,c", ","))      // [a,b,c]
	fmt.Println(strings.TrimSpace("   hello   ")) // hello

	// strconv 类型转换
	num, _ := strconv.Atoi(("123")) // string -> int
	fmt.Println(num + 1)            // 124

	str := strconv.Itoa(456) // int -> string
	fmt.Println(str + "元")   // 456元
}
