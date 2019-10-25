package main
import(
	"fmt"
	"net"
	"os"
)
func main() {
	name, err := os.Hostname()
	checkError(err)
	ipAddr, err := net.ResolveIPAddr("ip4", name)
	checkError(err)
	fmt.Println(ipAddr)
	conn, err := net.ListenIP("ip4:ip", ipAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.IPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("Receive from client", addr.String(), string(buf[0:n]))
	conn.WriteToIP([]byte("Welcome Client!"), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}