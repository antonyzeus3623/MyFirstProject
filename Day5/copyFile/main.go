package main

import (
	"fmt"
	"io"
	"os"
)

//借助io.Copy()实现一个拷贝文件函数。

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", srcName, err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dstName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) // 调用io.Copy()拷贝内容

}
func main() {
	_, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Printf("copy file failed, err:%v\n", err)
		return
	}
	fmt.Println("copy done!")

}
