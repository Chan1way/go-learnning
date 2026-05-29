package main

import (
	"errors"
	"fmt"
)

// 方式1： 用errors.New 创建错误
func getScore(name string, scores map[string]int) (int, error) {
	score, ok := scores[name]
	if !ok {
		return 0, errors.New("学生不存在: ")
	}
	return score, nil
}

// 方式2： 用fmt.Errorf 创建带格式的错误（更常用）
func checkScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("分数不合法: %d, 必须在0~100之间", score)
	}
	return nil
}

// func main() {
// 	scores := map[string]int{"小明": 95, "小红": 88}

// 	// 正常情况
// 	score, err := getScore("小明", scores)
// 	if err != nil {
// 		fmt.Println("出错：", err)
// 	} else {
// 		fmt.Println("小明的分数: ", score)
// 	}

// 	// 出错情况
// 	score, err = getScore("小王", scores)
// 	if err != nil {
// 		fmt.Println("出错", err)
// 	} else {
// 		fmt.Println("小王的分数: ", score)
// 	}

// 	err = checkScore(150)
// 	if err != nil {
// 		fmt.Println("出错：", err)
// 	}
// }
