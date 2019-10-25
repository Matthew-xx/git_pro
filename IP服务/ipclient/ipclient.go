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
	lAddr, err := net.ResolveIPAddr("ip4", service)
	checkError(err)
	name, err := os.Hostname()
	checkError(err)
	rAddr, err := net.ResolveIPAddr("ip4", name)
	checkError(err)
	conn, err := net.DialIP("ip4:ip", lAddr, rAddr)
	checkError(err)
	_, err = conn.WriteToIP([]byte("Hello Server!"), rAddr)
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[0:])
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