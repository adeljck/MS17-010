package utils

import (
	"fmt"
	"net"
)

func CheckDefaultRDPPort(ip, port string) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	defer conn.Close()
	if err != nil {
		// 连接失败，端口可能关闭
		return false
	}
	// 连接成功，关闭连接
	return true
}
