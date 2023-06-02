package output

import (
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

// GetIPList 获取本机可用ip列表
func GetIPList() (ipList []string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i2 := range interfaces {
		addrees, err := i2.Addrs()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, addr := range addrees {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ip4 := ipNet.IP.To4()
			if ip4 == nil {
				continue
			}
			ipList = append(ipList, ip4.String())
		}
	}
	return
}
