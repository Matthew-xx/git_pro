package main
import(
	"fmt"
	"net"
	"os"
)
func main() {
	service := ":5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)   //创建 Goroutine
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()          //逆序调用 Close() 保证连接能正常关闭
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client"))
		if err2 != nil {
			return
		}
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}