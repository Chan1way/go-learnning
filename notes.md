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

## Day 4 p
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

---

## Day 5 学习总结

  ### 错误处理

  Go 没有 try/catch，用返回值传递错误：
<pre>
  // 创建错误的两种方式
  errors.New("错误信息")           // 简单错误
  fmt.Errorf("错误: %d", value)   // 带格式的错误（更常用）

  // 调用方每次都要判断
  val, err := someFunc()
  if err != nil {
      fmt.Println("出错:", err)
  }
</pre>
- 出错时返回零值占位，正常时返回 nil 表示没有错误。


 ### 多文件拆分

同一个 package main 可以拆成多个文件：

 - main.go       → 只放 main 函数（程序入口）
 - student.go    → 放 Student 相关的类型和函数
 - errors_demo.go → 放其他功能函数

运行多文件项目：
go run *.go   # * 代表所有 .go 文件

---

## Day 6 学习总结

  **strings 字符串处理** 

  <pre>
  strings.ToUpper("hello")          // HELLO
  strings.Contains("hello", "ell")  // true
  strings.Split("a,b,c", ",")       // [a b c]
  strings.TrimSpace("  hello  ")    // hello
  </pre>
 

  **strconv 类型转换**
  <pre>
  num, _ := strconv.Atoi("123")  // string → int
  str := strconv.Itoa(456)       // int → string
  </pre>
  Go 强类型，不同类型不能直接拼接，必须先转换。

  **JSON 序列化 / 反序列化**
   
<pre>
  // 序列化：struct → JSON
  data, err := json.Marshal(s)
  fmt.Println(string(data))  // {"Name":"小明","Score":95}

  // 反序列化：JSON → struct
  json.Unmarshal([]byte(jsonStr), &s2)
</pre>

  **第一个 HTTP 服务器**

<pre>
  // 注册路由
  http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintln(w, "Hello, World!")
  })
  
  // 启动服务器
  http.ListenAndServe(":8080", nil)
</pre>
  
---

## Day7 学习总结
**Gin 基础**

- r := gin.Default()      // 创建路由器
- r.GET("/path", handler) // 注册路由
- r.Run(":8080")          // 启动服务器

gin.Default() 自带日志和错误恢复，是最常用的创建方式。

**gin.Context**
Gin 把请求和响应合并成一个 c *gin.Context：

**关键代码模式**
<pre>
  // 取路径参数并转 int
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil || id < 0 || id >= len(studentList) {
      c.JSON(http.StatusBadRequest, gin.H{"error": "id 不合法"})
      return  // 出错必须 return
  }

  // 解析请求体
  var s Student
  if err := c.ShouldBindJSON(&s); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
      return
  }

  // 删除 slice 元素
  studentList = append(studentList[:id], studentList[id+1:]...)
  // ...是展开操作符，主要把第2个 slice 展开成一个个元素传给 append
</pre>

## Day8学习总结
  **GORM 核心概念**
<pre>
  // 连接数据库
  DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{})

  // 自动建表
  DB.AutoMigrate(&Student{})
</pre>

**gorm.Model**
<pre>
  type Student struct {
      gorm.Model        // 自动加 4 个字段
      Name  string `json:"name"`
      Score int    `json:"score"`
  }
</pre>

**GORM 四个核心操作**
<pre>
  // 查所有
  var students []Student
  DB.Find(&students)

  // 查单个
  var student Student
  DB.First(&student, id)

  // 新增
  DB.Create(&student)
  
  // 修改
  DB.Save(&student)

  // 删除
  DB.Delete(&Student{}, id)
</pre>
  
  **错误检查**
  <pre>
  result := DB.Create(&student)
  if result.Error != nil {
      c.JSON(500, gin.H{"error": result.Error.Error()})
      return
  }

  GORM 操作默认不报错，要主动检查 result.Error。
  </pre>

## Day 9 学习总结：项目结构重组

### 项目分层结构
```
  go-learnning/
  ├── main.go          # 入口，只负责启动
  ├── router/
  │   └── router.go    # 路由注册
  ├── handler/
  │   └── student.go   # 接口处理函数
  ├── model/
  │   └── student.go   # 数据模型
  └── db/
      └── db.go        # 数据库连接
```
### 各层职责

  | 层 | 职责 | 类比 |
  |---|------|------|
  | model | 定义数据结构，不写逻辑 | 快递单模板 |
  | db | 建立数据库连接，提供全局 DB | 自来水公司铺水管 |
  | handler | 接收请求、处理业务、返回响应 | 厨师接单做菜出餐 |
  | router | 把 URL 和 handler 绑定 | 前台引导客人 |
  | main.go | 串起所有层，启动服务 | 公司开门营业 |

### 依赖关系
  ```
  main.go
    ├── db（初始化）
    └── router（注册路由）
          └── handler（处理函数）
                ├── db（操作数据库）
                └── model（数据结构）
  model 不依赖任何层（最底层）
  db 只依赖 model
  规则：依赖只能向下，不能反向
  ```

### 项目分层详解
  ```
请求进来
     ↓
  router    → 决定哪个 handler 处理
     ↓
  handler   → 接收请求、调数据库、返回响应
     ↓
    db        → 提供数据库连接
     ↓
  model     → 定义数据结构
  ```

#### model层（数据模型）
  **职责**：只负责定义数据长什么样，不做任何逻辑。
  ```
// model/student.go
  type Student struct {
      gorm.Model
      Name  string `json:"name"`
      Score int    `json:"score"`
  }
  ```
**类比**：快递单的格式模板，规定了有哪些字段。

**使用规则：**
  - 只放 struct 定义和字段
  - 不写任何业务逻辑
  - 其他层用 model.Student 来引用

#### db 层（数据库连接）
**职责**：只负责建立和维护数据库连接，提供全局 DB 变量给其他层用。

```
  // db/db.go
  var DB *gorm.DB   // 全局连接，整个项目共用

  func Init() {
      DB, _ = gorm.Open(...)
      DB.AutoMigrate(&model.Student{})
  }
```
  **类比**：自来水公司铺好水管，其他人直接拧开关用水，不用关心水怎么来的。

  **使用规则**：
  - 只放连接初始化代码
  - 其他层通过 db.DB.Find()、db.DB.Create() 等使用

#### handler 层（业务处理）
**职责**：接收请求 → 处理业务逻辑 → 返回响应。这是最核心的一层，大部分代码都在这里。
```
  // handler/student.go
  func GetStudents(c *gin.Context) {
      // 1. 从数据库取数据
      var students []model.Student
      db.DB.Find(&students)

      // 2. 返回响应
      c.JSON(http.StatusOK, students)
  }

  func CreateStudent(c *gin.Context) {
      // 1. 解析请求参数
      var student model.Student
      c.ShouldBindJSON(&student)

      // 2. 写入数据库
      db.DB.Create(&student)

      // 3. 返回响应
      c.JSON(http.StatusOK, student)
  }
```
  **类比**：餐厅厨师，接到点单（请求）→ 做菜（处理逻辑）→ 出餐（返回响应）。

  **使用规则**：
  - 每个函数对应一个接口
  - 函数签名固定：func Xxx(c *gin.Context)
  - 不要在这里写 SQL，用 GORM 的方法操作数据库

#### router 层（路由注册）
  
  **职责**：只负责把 URL 路径和 handler 函数绑定在一起。

```
  // router/router.go
  func Setup() *gin.Engine {
      r := gin.Default()

      students := r.Group("/students")  // 路由分组
      {
          students.GET("", handler.GetStudents)       // GET /students
          students.GET("/:id", handler.GetStudent)    // GET /students/1
          students.POST("", handler.CreateStudent)    // POST /students
          students.PUT("/:id", handler.UpdateStudent) // PUT /students/1
          students.DELETE("/:id", handler.DeleteStudent) // DELETE /students/1
      }

      return r
  }
```
  **类比**：前台接待，根据客人需求（URL）引导去找对应的工作人员（handler）。

  **使用规则**:
  - 只放路由注册，不写业务逻辑
  - 用 r.Group 给相关路由分组，保持整洁
#### main.go（程序入口）
  
**职责**：只负责把所有层串起来，启动服务。

```
  func main() {
      db.Init()        // 1. 初始化数据库
      r := router.Setup() // 2. 注册路由
      r.Run(":8080")   // 3. 启动服务器
  }
```
**类比**：公司开门营业，依次通水（db）、安排前台（router）、开门迎客（Run）。

  ### 新增接口只需两步
  **第一步：handler 写处理函数**
  ```go
  func GetStudentsByScore(c *gin.Context) {
      score := c.Query("score")
      var students []model.Student
      db.DB.Where("score >= ?", score).Find(&students)
      c.JSON(http.StatusOK, students)
  }

  第二步：router 注册路由
  students.GET("/search", handler.GetStudentsByScore)

  今天踩的坑

  ┌──────────────────┬────────────────────────────┬────────────────────────────────┐
  │       问题        │            原因             │              解决              │
  ├──────────────────┼────────────────────────────┼────────────────────────────────┤
  │ import 路径报错    │ 模块名和 import 路径不一致    │ 检查 go.mod 第一行，统一模块名    │
  ├──────────────────┼────────────────────────────┼────────────────────────────────┤
  │ 端口 8080 被占用   │ 旧服务器进程没关              │ lsof -ti:8080 | xargs kill -9  │
  └──────────────────┴────────────────────────────┴────────────────────────────────┘

  关键命令
  
  # 查看模块名
  head -1 go.mod

  # 释放被占用的端口
  lsof -ti:8080 | xargs kill -9

  # 启动服务器
  go run main.go

  ```

## Day10 中间件 Logger() Recovery() 、统一格式
#### 自定义中间件
  - 中间件是套在 handler 外面的一层，每个请求都会经过
  - c.Next() = 执行后续 handler，跑完再回来
  - r.Use() = 注册中间件

#### Recovery 中间件
  - defer + recover() 捕获 panic，防止服务器崩溃
  - c.Abort() = 终止后续所有执行
  
#### 统一响应格式
  - 定义 Response 结构体，所有接口返回 code、message、data
  - Success() 和 Fail() 两个辅助函数，handler 里不再散写 c.JSON()

#### 提问
  ##### \* 是什么？
  出现在类型前面表示指针类型，存的是变量的内存地址（门牌号），不是值本身。

  ##### &变量名 是什么？
  取这个变量的地址。* 和 & 是一对：*类型 声明这是个地址类型，&变量 拿到某个变量的地址。

  ##### r.Use() 是做什么的？
  注册中间件，让每个请求在进入 handler 之前/之后都经过指定的函数。

  ##### c.Next() 为什么叫"把控制权交给后面的 handler"？
  Gin 把中间件和 handler 排成一个队列，c.Next() 就是"先去执行队列里的下一个，跑完了回来继续"。

  ##### interface{} 是什么？
  表示任意类型，当一个变量可能是不同类型时使用（比如 data 有时是单个学生，有时是列表，有时是 nil）。

  ##### Go 时间格式为什么是 2006-01-02？
  Go 不用 YYYY-MM-DD，而是用一个固定参考时间来表示格式，年必须写 2006，写成别的年份就会出错