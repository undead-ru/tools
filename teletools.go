package teletools

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

func DecimalToBase(dec, rad int) (res string) {

	myascii := "0123456789abcdefghijklmnopqrstuvwxyz" // max base 36

	if rad <= 1 || rad > len(myascii) || dec <= 0 {
		return "0"
	}
	for i := dec; i > 0; i = i / rad {
		res = fmt.Sprintf("%s%s", myascii[(i%rad):(i%rad)+1], res)
	}

	return
}

var ReSplitByLen = regexp.MustCompile("([\n\r]|[\r\n])")

func SplitByLen(strIn string, strLen int) []string {

	var strN []string
	var str2 string

	str := ReSplitByLen.ReplaceAllString(strIn, "\n")
	for _, str := range strings.Split(str, "\n") {
		if len(str2)+len(str) >= strLen {
			strN = append(strN, str2)
			str2 = str + "\n"
		} else {
			str2 += str + "\n"
		}
	}

	strN = append(strN, str2)

	return strN
}

func GetIPs() (listIPs []string) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Can't get Interfaces: %v", err)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				listIPs = append(listIPs, string(ipnet.IP.String()))
			}
		}
	}

	return
}
