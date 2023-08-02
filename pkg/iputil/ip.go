package util

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"net"
)

func GetRemoteRealIp(c *gin.Context) string {
	remoteIp := exnet.ClientPublicIP(c.Request)
	if remoteIp == "" {
		remoteIp = exnet.ClientIP(c.Request)
	}
	if remoteIp == "" {
		remoteIp = c.ClientIP()
	}
	return remoteIp
}

func IsContainsIp(cidr string, ip net.IP) (bool, error) {
	if cidr == "" {
		return true, nil
	}
	if pubIp := net.ParseIP(cidr); pubIp != nil {
		if pubIp.Equal(ip) {
			return true, nil
		}
	} else {
		_, pubCidr, err := net.ParseCIDR(cidr)
		if err != nil {
			return false, err
		}
		if pubCidr.Contains(ip) {
			return true, nil
		}
	}
	return false, nil
}

func IsPublicIP(IP net.IP) bool {
	if IP == nil {
		return false
	}
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		case ip4[0] == 100 && ip4[1] >= 64:
			return false
		default:
			return true
		}
	}
	return false
}
