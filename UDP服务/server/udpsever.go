package main
import(
	"fmt"
	"net"
	"os"
)
func main() {
	service := ":5001"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("Receive from client", addr.String(), string(buf[0:n]))
	conn.WriteToUDP([]byte("Welcome Client!"), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}