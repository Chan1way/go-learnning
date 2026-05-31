package main

import "fmt"

// type Student struct {
// 	Name  string
// 	Score int
// }

func printStudents(students []Student) {
	for _, student := range students {
		fmt.Println(student.Name, "的成绩是：", student.Score)
	}
}

func countStats(students []Student) map[string]int {
	stats := map[string]int{"及格": 0, "不及格": 0}
	for _, student := range students {
		if student.Score >= 60 {
			stats["及格"]++
		} else {
			stats["不及格"]++
		}
	}
	return stats
}

// ---
// 练习题（20分钟）

// 在 student.go 里新增一个函数：

// 函数名：getTopStudent
// 传入：[]Student
// 返回：分数最高的那个 Student 和 error
// 要求：如果 slice 是空的，返回错误 "没有学生数据"

// 在 main.go 里调用它，打印最高分学生的名字和分数。

// 自己先写，写完贴给我。

func getTopStudent(students []Student) (Student, error) {
	if students == nil {
		return Student{}, fmt.Errorf("没有学生数据")
	} else {
		max_student := students[0]
		for _, student := range students {
			if student.Score > max_student.Score {
				max_student = student
			}
		}
		return max_student, nil
	}
}
