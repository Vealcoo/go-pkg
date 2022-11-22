package hash

import (
	"crypto/md5"
	"fmt"

	"hash/crc32"
)

var chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func StringHash(in string) string {
	h := crc32.NewIEEE()
	h.Write([]byte(in))

	return encode(h.Sum32())
}

func StringMD5(in string) string {
	has := md5.Sum([]byte(in))
	md5str := fmt.Sprintf("%x", has)

	return md5str
}

func encode(num uint32) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
