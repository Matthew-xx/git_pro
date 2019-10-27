package main
//打开一个存在的文件，将原来的内容读出来，显示在终端，并且追加 5 句：hello，C语言中文网
import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main() {
	filePath := "f:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("open file err",err)
	}
	//及时关闭file句柄
	defer file.Close()
	//读原来文件的内容，并且显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString("hello，C语言中文网。 \r\n")
	}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}
