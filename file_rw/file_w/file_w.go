package main
//创建一个新文件，写入 5 句：hello,Go语言
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	//创建一个新文件，写入内容 5句 hello,Go语言
	filePath := "f:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil{
		fmt.Println("open file err",err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("hello,Go语言 \r\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}