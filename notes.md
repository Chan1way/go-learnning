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

### 第四步：第一个 Go 程序（40分钟）
  
1.在同一个目录里：

2.初始化 Go 项目
  go mod init go-learning
  
3.在 VS Code 里新建文件 main.go，写入：

  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, World!")
      fmt.Println("我开始学Go了")
  }

  回终端跑：
  go run main.go
  
  看到输出后，把代码推到 GitHub：
  git add .
  git commit -m "add hello world"
  git push
