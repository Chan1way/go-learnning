# 我的go学习记录
## day2
### 第一步：配置git
git config --global user.name "你的名字"
git config --global user.email "邮箱"

验证：
git config --list

### 第二步：在github 创建第一个仓库
1. 打开 https://github.com，没账号先注册
2. 点右上角 + → New repository
3. 填写：
 - Repository name：go-learning
 - 勾选 Add a README file
 - 点 Create repository
4. 点绿色 Code 按钮，复制 HTTPS 地址（类似 https://github.com/你的名字/go-learning.git）

### 第三步：Git 5 个命令实战
1. 把仓库拉到本地
git clone https://github.com/你的名字/go-learning.git

2. 进入项目目录
  cd go-learning
  
3. 看状态（目前是干净的）
  git status

   创建一个文件：
  echo "# 我的Go学习记录" > notes.md

4. 看状态（会显示 notes.md 是新文件）
  git status

5. 标记要提交
  git add notes.md
  把 notes.md 放进暂存区（告诉 Git：我要提交这个文件）。类似于：打包行李放进箱子，但还没发货。

6. 提交
  git commit -m "add notes"
  把暂存区的内容提交到本地仓库，-m 后面是提交说明。类似于：箱子封好了，贴上标签"add notes"，但还在你手上。

7. 推到 GitHub
  git push
  把本地仓库的提交推到 GitHub。类似于：把箱子发出去，GitHub 上就能看到了。

8. 去 GitHub 刷新页确认文件是否更新

### 第四步：第一个 Go 程序
  
1. 在同一个目录里：

2. 初始化 Go 项目
  go mod init go-learning
  
3. 在 VS Code 里新建文件 main.go，写入：
<pre>
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, World!")
      fmt.Println("我开始学Go了")
  }
</pre>
 4. 回终端跑：
  go run main.go
  
5. 看到输出后，把代码推到 GitHub：
  git add .
  git commit -m "add hello world"
  git push

---

## day3
### 变量
  <pre>
  var name string = "小明"  // 完整写法`
  age := 25                // 简短写法（最常用）
  := 自动判断类型，是 Go 最常用的变量声明方式。
  </pre>

### 函数
<pre>
  func add(a int, b int) int {
      return a + b
  }
</pre>
  - 参数：名字在前，类型在后
  - 返回值类型写在括号后面
  - 有返回值声明就必须有 return
 

 **Go 特色：多返回值 + 错误处理**
  <pre>
  func divide(a int, b int) (int, error) {
      if b == 0 {
          return 0, fmt.Errorf("除数不能为0")
      }
      return a / b, nil
  }

  val, err := divide(10, 2)
  if err != nil {
      fmt.Println("出错了:", err)
  }
  </pre>

  - 函数可以返回多个值
  - nil 表示没有错误
  - err != nil 是 Go 代码里最常见的模式

### 循环
<pre>
  // 普通循环，Go 只有 for 一种循环，没有 while。
  for i := 1; i <= 5; i++ { }

  // 遍历数组（最常用）
  for _, value := range slice {
      // _ 忽略索引，value 是实际的值
  }
</pre>
---

## Day 4 
### struct（结构体）
<pre>
  type Student struct {
      Name  string
      Score int
  }

  s1 := Student{Name: "小明", Score: 95}
  fmt.Println(s1.Name, s1.Score)  // 用 . 访问字段

  Go 没有 class，用 struct 表示一个对象。
</pre>
  
  ### slice（切片）
<pre>
  students := []Student{s1, s2, s3}  // 创建
  students = append(students, s4)     // 追加
  students[1:3]                       // 截取

  for _, student := range students {  // 遍历
      fmt.Println(student.Name)
  }
  
  动态数组，append 追加元素，range 遍历。
</pre>

###  map（字典）
<pre>
  // 声明：map[键类型]值类型
  scores := map[string]int{
      "小明": 95,
  }

  scores["小李"] = 90        // 新增/修改
  scores["小明"]             // 读取

  // 判断 key 是否存在
  score, ok := scores["小王"]
  if !ok {
      fmt.Println("不存在")
  }
  
  for name, score := range scores {  // 遍历
      fmt.Println(name, score)
  }
  </pre>