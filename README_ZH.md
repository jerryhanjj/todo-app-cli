# Todo App CLI

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

**Language:** 中文 | [English](README.md)

一个简单而强大的命令行待办事项应用程序，使用Go语言编写。

## ✨ 特性

- 📝 添加新的待办事项
- 📋 查看所有待办事项
- ✅ 标记待办事项为完成
- 🗑️ 删除待办事项
- 💾 使用JSON格式持久化存储
- 🔢 智能序号管理（连续显示序号）
- 🛡️ 安全的ID生成机制
- 🎯 基于Cobra的现代化子命令界面
- 🔄 向后兼容标志模式
- 📚 自动生成帮助文档
- 🎨 丰富的命令别名支持

## 📦 安装方法

### 方法1：使用Go Install安装（推荐）

**前提条件：** 需要安装Go 1.19或更高版本

```bash
# 安装最新版本
go install github.com/jerryhanjj/todo-app-cli@latest

# 或者安装特定版本（例如v1.0.0）
go install github.com/jerryhanjj/todo-app-cli@v1.0.0

# 运行应用
todo-app-cli list
```

### 方法2：从源码安装

**前提条件：** 需要安装Go 1.19或更高版本

```bash
# 1. 克隆仓库
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# 2. 构建可执行文件
go build -o todo main.go

# 3. 运行应用
./todo list
```

### 方法3：直接运行（开发模式）

```bash
# 克隆仓库
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# 直接运行（每次都会重新编译）
go run main.go list
```

### 方法4：使用Makefile（推荐给开发者）

```bash
# 克隆仓库
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# 使用Makefile构建
make build

# 运行应用
./todo list

# 或者直接运行（开发模式）
make run ARGS="list"
make run ARGS="add '新任务'"
```

### 方法5：安装到系统PATH

```bash
# 在项目目录下执行
go build -o todo main.go

# 移动到系统PATH目录（可选）
sudo mv todo /usr/local/bin/

# 或者使用Makefile
make install

# 现在可以在任何地方使用
todo list
```

## 🚀 快速开始

```bash
# 查看帮助信息
./todo --help

# 新的子命令方式（推荐）
./todo add "学习Go编程"
./todo list
./todo complete 1
./todo delete 2

# 兼容旧版本的标志方式
./todo -a "学习Go编程"
./todo -l
./todo -c 1
./todo -d 2
```

## 📖 详细用法

### 两种使用方式

本工具支持两种使用方式，您可以根据个人喜好选择：

#### 方式1: 子命令方式（推荐）
```bash
./todo add "买菜"          # 添加待办事项
./todo list                # 查看列表
./todo complete 1          # 完成第1项
./todo delete 2            # 删除第2项
```

#### 方式2: 标志方式（兼容旧版本）
```bash
./todo -a "买菜"           # 添加待办事项
./todo -l                  # 查看列表
./todo -c 1                # 完成第1项
./todo -d 2                # 删除第2项
```

### 查看帮助
```bash
# 查看主命令帮助
./todo --help

# 查看子命令帮助
./todo add --help
./todo list --help
./todo complete --help
./todo delete --help
```

### 添加待办事项
```bash
# 子命令方式
./todo add "买菜"
./todo add "写代码"
./todo add "锻炼身体"

# 标志方式
./todo -a "买菜"
./todo --add "写代码"
```

### 查看待办事项列表
```bash
### 查看待办事项列表
```bash
# 子命令方式
./todo list
# 或使用别名
./todo ls
./todo l

# 标志方式
./todo -l
./todo --list
```
输出示例:
```
待办事项:
[✓] 1: 学习Go编程
[ ] 2: 买菜  
[ ] 3: 写代码
```

### 完成待办事项
```bash
# 子命令方式
./todo complete 2
# 或使用别名
./todo done 2
./todo c 2

# 标志方式
./todo -c 2
./todo --complete 2
```

### 删除待办事项
```bash
# 子命令方式
./todo delete 3
# 或使用别名
./todo del 3
./todo d 3
./todo rm 3

# 标志方式
./todo -d 3
./todo --delete 3
```
```
输出示例：
```
待办事项:
[✓] 1: 学习Go编程
[ ] 2: 买菜  
[ ] 3: 写代码
```

### 完成待办事项
```bash
# 使用序号完成待办事项
./todo complete 2
# 或使用别名
./todo done 2
./todo c 2
```

### 删除待办事项
```bash
# 使用序号删除待办事项
./todo delete 3
# 或使用别名
./todo del 3
./todo d 3
./todo rm 3
```

## 🎯 命令概览

### 子命令方式
| 命令 | 别名 | 描述 | 示例 |
|------|------|------|------|
| `add` | - | 添加新的待办事项 | `todo add "买菜"` |
| `list` | `ls`, `l` | 显示所有待办事项 | `todo list` |
| `complete` | `done`, `c` | 标记待办事项为已完成 | `todo complete 1` |
| `delete` | `del`, `d`, `rm` | 删除待办事项 | `todo delete 2` |
| `help` | - | 显示帮助信息 | `todo help` |

### 标志方式（兼容旧版本）
| 短标志 | 长标志 | 描述 | 示例 |
|--------|--------|------|------|
| `-a` | `--add` | 添加新的待办事项 | `todo -a "买菜"` |
| `-l` | `--list` | 显示所有待办事项 | `todo -l` |
| `-c` | `--complete` | 标记待办事项为已完成 | `todo -c 1` |
| `-d` | `--delete` | 删除待办事项 | `todo -d 2` |
| `-h` | `--help` | 显示帮助信息 | `todo -h` |

## 📁 项目结构

```
todo-app-cli/
├── main.go                 # 应用程序入口点
├── cmd/                    # 命令定义（Cobra框架）
│   ├── root.go             # 根命令和共享功能
│   ├── add.go              # 添加命令实现
│   ├── list.go             # 列表命令实现
│   ├── complete.go         # 完成命令实现
│   └── delete.go           # 删除命令实现
├── internal/
│   └── todo/
│       └── todo.go          # 待办事项逻辑和数据结构
├── data/
│   └── todos.json           # JSON格式的持久化存储
├── Makefile                 # 构建和开发工具
├── go.mod                   # Go模块文件
├── go.sum                   # Go依赖校验文件
├── LICENSE                  # 许可证文件
├── README_ZH.md             # 项目说明文档（中文）
└── README.md                # 项目说明文档（英文）
```

## 🔧 开发环境设置

### 前提条件
- Go 1.19+ ([安装指南](https://golang.org/doc/install))
- Git

### 本地开发
```bash
# 1. 克隆项目
git clone https://github.com/jerryhanjj/todo-app-cli.git
cd todo-app-cli

# 2. 安装依赖（如果有的话）
go mod tidy

# 3. 使用Makefile进行开发
make run ARGS="list"          # 运行列表命令
make run ARGS="add '测试'"    # 添加待办事项
make build                    # 构建可执行文件
make test                     # 运行测试
make fmt                      # 格式化代码
make clean                    # 清理构建文件

# 4. 或者传统方式
go run main.go list
go build -o todo main.go
```

### Makefile命令
```bash
make help          # 查看所有可用命令
make build         # 构建应用
make run ARGS="..."# 运行开发模式
make test          # 运行测试
make clean         # 清理构建文件
make install       # 安装到系统PATH
make fmt           # 格式化代码
make vet           # 代码检查
```

## 💡 特性详解

### 智能序号管理
- 显示连续的序号（1, 2, 3...），即使删除了中间的项目
- 用户操作使用序号，内部维护唯一ID
- 避免了序号跳跃的困扰

### 数据持久化
- 所有数据保存在`data/todos.json`文件中
- 自动创建数据目录和文件
- JSON格式便于查看和调试

### 错误处理
- 完善的边界检查
- 友好的中文错误提示
- 防止无效操作

## 🤝 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork这个仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启一个Pull Request

## 📄 许可证

本项目采用MIT许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 🔮 计划功能

- [ ] 添加待办事项优先级
- [ ] 支持待办事项分类/标签
- [ ] 添加截止日期功能
- [ ] 导出功能（JSON, CSV, TXT等）
- [ ] 搜索和过滤功能
- [ ] 彩色输出
- [ ] 命令行自动补全
- [ ] 配置文件支持

## 📝 更新日志

### v2.0.0 (当前版本)
- ✅ **重大升级**: 从pflag迁移到Cobra框架
- ✅ 现代化子命令界面（`todo add`, `todo list` 等）
- ✅ **向后兼容**: 仍支持旧版标志（`-a`, `-l`, `-c`, `-d`）
- ✅ 命令别名支持（`ls`, `done`, `del` 等）
- ✅ 自动生成帮助文档
- ✅ 改进的代码组织结构，独立命令文件
- ✅ 增强的用户体验和错误处理
- ✅ 为未来功能扩展奠定基础

### v1.0.0 (旧版本)
- ✅ 基础CRUD功能
- ✅ JSON持久化存储
- ✅ 智能序号管理
- ✅ 中文错误提示
- ✅ 基于pflag的命令行界面

## 🐛 问题反馈

如果你遇到任何问题或有建议，请：

1. [创建Issue](https://github.com/jerryhanjj/todo-app-cli/issues)
2. 邮件联系：jerryhanjj@126.com
3. 在GitHub上start ⭐ 这个项目

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者！

特别感谢：
- Go社区提供的优秀文档和工具
- 帮助改进这个项目的贡献者们

---

**Happy Coding! 🚀**