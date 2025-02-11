# HashUtil SDK

`HashUtil` 是一个简单的 Go SDK，提供了 MD5 和 SHA-256 哈希计算功能。

## 安装

使用以下命令安装 SDK：

```sh
go get github.com/yourusername/hashutil
```


## 项目目录结构

```bash
hashutil/
├── .git/                          # Git 管理
├── .github/workflows/             # CI 配置
│   └── ci.yml                     # GitHub Actions 工作流
├── examples/                      # 示例代码
│   └── example.go
├── internal/                      # 私有包
│   └── hashalgo.go                 # 私有哈希计算逻辑
├── pkg/                           # 公共包
│   └── hashutil.go                 # SDK 公开功能
├── test/                          # 单元测试
│   └── hashutil_test.go            # SDK 测试代码
├── go.mod                          # Go 模块定义文件
├── go.sum                          # Go 依赖文件
└── README.md                       # 项目文档
```

## 使用示例

```go
package main

import (
	"fmt"
	"github.com/yourusername/hashutil"
)

func main() {
	text := "hello"
	fmt.Println("MD5 Hash:", hashutil.MD5Hash(text))
	fmt.Println("SHA256 Hash:", hashutil.SHA256Hash(text))
}
```

## 运行测试

```
go test ./test/
```

##  总结

通过上述步骤，我们创建了一个标准的 Go SDK 项目，包含了以下部分：

- **核心功能 (`pkg/`)**：提供 MD5 和 SHA-256 哈希计算功能。 
- **私有功能 (`internal/`)**：用于内部逻辑，如 `CalculateMD5Hash`。 
- **单元测试 (`test/`)**：为 SDK 提供测试。
- **示例代码 (`examples/`)**：展示如何使用 SDK。
- **CI 配置 (`.github/workflows/`)**：通过 GitHub Actions 自动化运行测试
- **文档 (`README.md`)**：项目的基本介绍、安装和使用说明。 将 SDK 发布到 GitHub 后，其他开发者可以通过 `go get` 安装并使用。
