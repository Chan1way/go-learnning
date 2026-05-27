// day3 Go 语法核心——变量、函数、循环
// // 变量
// package main
// import "fmt"

// func main() {
// 	//写法1: 完整写法
// 	var name string = "小明"
// 	var age int = 25

// 	// 写法2: 简短写法（最常用）
// 	score := 90
// 	isStudent := true

// 	fmt.Println(name, age, score, isStudent)
// }

// // 函数
// package main

// import "fmt"
// 定义一个函数：传入两个数，返回它们的和
// func add(a int, b int) int {
// 	return a + b
// }

// Go 特色：可以返回多个值
// func divice(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("除数不能为0")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	result := add(3, 5)
// 	fmt.Println("3+5=", result)

// 	val, err := divice(10, 0)
// 	if err != nil {
// 		fmt.Println("出错了", err)
// 	} else {
// 		fmt.Println("10/2=", val)
// 	}
// }

// 循环
// package main

// import "fmt"

// func main() {
// 	// 普通循环
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("第", i, "次")
// 	}

// 	// 遍历数组
// 	fruits := []string{"苹果", "香蕉", "橙子"}
// 	for index, fruit := range fruits {
// 		fmt.Println(index, "->", fruit)
// 	}
// }

// 练习题
//   用上面学的知识，自己写一个函数，要求：

//   - 函数名：isPass
//   - 传入一个分数（int 类型）
//   - 分数 >= 60 返回 "通过"，否则返回 "不通过"
//   - 在 main 里分别测试 59、60、90 三个分数

package main

import "fmt"

func isPass(score int) string {
	if score >= 60 {
		return "通过"
	} else {
		return "不通过"
	}
}

func main() {
	scores := []int{59, 60, 90}
	for _, score := range scores {
		result := isPass(score)
		fmt.Println(score, "->", result)
	}

}
