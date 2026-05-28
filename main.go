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

// package main

// import "fmt"

// func isPass(score int) string {
// 	if score >= 60 {
// 		return "通过"
// 	} else {
// 		return "不通过"
// 	}
// }

// func main() {
// 	scores := []int{59, 60, 90}
// 	for _, score := range scores {
// 		result := isPass(score)
// 		fmt.Println(score, "->", result)
// 	}

// }

// 今天踩的两个坑（记住）

//   ┌────────────────────────────────────────┬──────────────────────────┬──────────────────────────────────┐
//   │                  错误                   │           原因            │               修复               │
//   ├────────────────────────────────────────┼──────────────────────────┼──────────────────────────────────┤
//   │ 函数声明了返回类型但用 fmt.Println         │ 声明了返回值就必须          │ 把 fmt.Println 改成 return        │
//   │ 输出                                    │ return                   │                                  │
//   ├────────────────────────────────────────┼──────────────────────────┼──────────────────────────────────┤
//   │ for score := range scores 拿到的是索引   │ range 默认返回索引         │ 改成 for _, score := range        │
//   │                                        │                          │ scores                           │
//   └────────────────────────────────────────┴──────────────────────────┴──────────────────────────────────┘

// day 4 : struct、slice、map
package main

import "fmt"

// struct
// 定义一个学生结构体
// Go 没有 class，用 struct 来表示一个对象：
// type Student struct {
// 	Name  string
// 	Age   int
// 	Score float64
// }

// func main() {
// 	// 创建一个学生
// 	s1 := Student{
// 		Name:  "小明",
// 		Age:   18,
// 		Score: 95.5,
// 	}

// 	// 访问字段
// 	fmt.Println(s1.Name, s1.Age, s1.Score)

// 	// 修改字段
// 	s1.Score = 98.0
// 	fmt.Print("修改后:", s1.Score)
// }

// slice（切片）
// func main() {
// 	// 创建 slice
// 	student := []string{"小明", "小红", "小张"}

// 	// 追加元素
// 	student = append(student, "小李")

// 	// 遍历
// 	for i, name := range student {
// 		fmt.Println(i, name)
// 	}

// 	// 截取（索引 1 到 3，不包含 3）
// 	fmt.Println(student[1:3])
// }

// map（字典）
// map就是键值对，类似python的dict

// func main() {
// 	// 创建map
// 	scores := map[string]int{
// 		"小明": 95,
// 		"小红": 88,
// 		"小张": 72,
// 	}

// 	// 读取
// 	fmt.Println("小明的分数：", scores["小明"])

// 	// 新增/修改
// 	scores["小李"] = 90

// 	// 判断key是否存在
// 	score, ok := scores["小王"]
// 	if !ok {
// 		fmt.Println("小王不存在")
// 	} else {
// 		fmt.Println("小王的分数：", score)
// 	}

// 	for name, score := range scores {
// 		fmt.Println(name, "->", score)
// 	}
// }

// 练习题：综合运用（30分钟）

// 用 struct + slice + map 写一个简单的学生管理程序：

// 要求：
// 1. 定义 Student struct，包含 Name、Score 两个字段
// 2. 创建一个存放学生的 slice，包含 3 个学生
// 3. 遍历 slice，打印每个学生的名字和成绩
// 4. 用 map 统计：60分以下多少人，60分以上多少人

// 自己先写，写不出来再问我。写完贴给我看。

type Student struct {
	Name  string
	Score int
}

func main() {
	s1 := Student{
		Name:  "小陈",
		Score: 100,
	}

	s2 := Student{
		Name:  "小李",
		Score: 60,
	}

	s3 := Student{
		Name:  "小沈",
		Score: 50,
	}

	students := []Student{s1, s2, s3}
	for _, student := range students {
		fmt.Println(student.Name, "的成绩是：", student.Score)
	}

	stats := map[string]int{
		"及格":  0,
		"不及格": 0,
	}
	for _, student := range students {
		if student.Score >= 60 {
			stats["及格"]++
		} else {
			stats["不及格"]++
		}
	}
	for stat, count := range stats {
		fmt.Println(stat, "->", count)
	}

}
