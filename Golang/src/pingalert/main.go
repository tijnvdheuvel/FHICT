
package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"time"
)

//Okay de adapter is blijkbaar niet vindbaar maar geen idee waarom dit niet werkt....
var adapter string = "eth0" //Adapter naam (Voor een lijst PowShell [Get-NetAdapter -Name "*"]
var snaplen int32 = 65535 //Lengte van de bytes van de hele packet
var promisc bool = false
var err error
var timeout time.Duration = -1 * time.Second
var handle *pcap.Handle

func main(){

//PING maakt gebruik van ICMP protocl want functioneert op poort 1, TCP
handle, err = pcap.OpenLive(adapter,snaplen,promisc,timeout)
if err != nil{
	panic (err)
}
defer handle.Close()

var filter string = "dst port 1"
err = handle.SetBPFFilter(filter)
if err != nil {
	panic(err)
}
packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
for packet := range packetSource.Packets() {
	fmt.Println("Incomming ping....")
	fmt.Println("================")
	fmt.Println(packet)

}
}



 /*
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":1")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			fmt.Println(c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
*/