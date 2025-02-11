## 01.创建SDK包

### **步骤 1：创建 SDK 包**

我们创建一个名为 `hashutil` 的 Go SDK，它包含哈希计算功能。

#### **1.1 初始化 Go 模块**

```bash
mkdir hashutil
cd hashutil
go mod init github.com/Tom7481079/hashutil
```

------

#### **1.2 创建 SDK 代码**

在 `hashutil` 目录下创建 `hashutil.go`

```go
package hashutil

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5Hash 计算字符串的 MD5 哈希值
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA256Hash 计算字符串的 SHA-256 哈希值
func SHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
```

------

#### **1.3 编写单元测试**

在 `hashutil` 目录下创建 `hashutil_test.go`：

```go
package hashutil

import (
	"testing"
)

func TestMD5Hash(t *testing.T) {
	result := MD5Hash("hello")
	expected := "5d41402abc4b2a76b9719d911017c592"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestSHA256Hash(t *testing.T) {
	result := SHA256Hash("hello")
	expected := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
```

------

#### **1.4 发布到 GitHub**

```bash
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/Tom7481079/hashutil.git
git push -u origin main
```

------

### **步骤 2：在其他项目中使用 SDK**

在另一个 Go 项目中使用该 SDK，可以这样做：

#### **2.1 安装 SDK**

```bash
go get github.com/Tom7481079/hashutil
```

#### **2.2 在代码中使用**

```go
package main

import (
	"fmt"
	"github.com/Tom7481079/hashutil"
)

func main() {
	text := "hello"
	fmt.Println("MD5:", hashutil.MD5Hash(text))
	fmt.Println("SHA256:", hashutil.SHA256Hash(text))
}
```