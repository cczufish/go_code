package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data := []byte("this is test, hello world, keep coding")
	fmt.Printf("%x \n", sha1.Sum(data))

	h := sha1.New()
	io.WriteString(h, "this is test, hello world, keep coding")
	fmt.Printf("%x \n", h.Sum(nil))

	fmt.Printf("%x \n", shaFile("./file.text"))
}

//shaFile利用sha1算法将目标文件生成哈希值
func shaFile(filePath string) []byte {
	f, err := os.Open("file.text")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}
