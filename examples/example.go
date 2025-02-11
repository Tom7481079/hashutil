package main

import (
	"fmt"
	"github.com/Tom7481079/hashutil/pkg"
)

func main() {
	text := "hello"
	fmt.Println("MD5 Hash:", hashutil.MD5Hash(text))
	fmt.Println("SHA256 Hash:", hashutil.SHA256Hash(text))
}
