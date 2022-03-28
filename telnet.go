package main

import (
	"github.com/reiver/go-telnet"
	"fmt"
	"strings"
)
func ReaderTelnet(conn *telnet.Conn, expect string) (out string) {
    var buffer [1]byte
    recvData := buffer[:]
    var n int
    var err error

    for {
        n, err = conn.Read(recvData)
        //fmt.Println("Bytes: ", n, "Data: ", recvData, string(recvData))
		//fmt.Print(string(recvData))
        if n <= 0 || err != nil || strings.Contains(out, expect) {
            break
        } else {
            out += string(recvData)
        }
		conn.Close()
    }
    return out
}
func ConnectTelnet(host string,komenda string)string{
	conn, err := telnet.DialTo(host+":8000")
	if err != nil {
        fmt.Println("dial error:", err)
        return host+" >> Error"
    }
    conn.Write([]byte(komenda))
    conn.Write([]byte("\n"))
	
	wynik := ReaderTelnet(conn,"\n\n")
	return "=================================== \n"+host+"\n"+wynik+"=================================== \n"
}

func main() {
	wynik := ConnectTelnet("FQDN","Telnet Command")
	fmt.Print(wynik)
}