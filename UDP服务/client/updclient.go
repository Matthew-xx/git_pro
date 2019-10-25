package main
import(
	"fmt"
	"net"
	"os"
)
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.Write([]byte("Hello server!"))
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	fmt.Println("Reply form server", addr.String(), string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}