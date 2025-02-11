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

