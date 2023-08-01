package encode

import (
	"crypto/md5"
	"encoding/hex"
	"hash/crc32"
	"strconv"
)

func CalMd5(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}

func CalStrListMd5(stringList []string) string {
	h := md5.New()
	for _, s2 := range stringList {
		h.Write([]byte(s2))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// Crc32HashCode hashes a string to a unique hashcode.
// crc32 returns an uint32, but for our use we need
// and non-negative integer. Here we cast to an integer
// and invert it if the result is negative.
func Crc32HashCode(b []byte) string {
	v := int(crc32.ChecksumIEEE(b))
	if v >= 0 {
		return strconv.Itoa(v)
	}
	if -v >= 0 {
		return strconv.Itoa(-v)
	}
	return ""
}
