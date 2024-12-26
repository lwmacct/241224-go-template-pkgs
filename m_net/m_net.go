package m_net

import (
	"crypto/rand"
	"fmt"
	"net"
)

func IsValidMAC(mac string) bool {
	_, err := net.ParseMAC(mac)
	return err == nil
}

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}

// 生成一个随机的本地管理单播地址
func GenerateRandomMAC() string {
	buf := make([]byte, 6)
	if _, err := rand.Read(buf); err != nil {
		panic("crypto/rand failed to generate random bytes: " + err.Error())
	}
	// 设置本地管理位（第 7 位）和单播位（第 8 位）
	buf[0] = (buf[0] & 0xfe) | 0x02
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}
