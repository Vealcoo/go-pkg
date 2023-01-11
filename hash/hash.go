package hash

import (
	"crypto/md5"
	"fmt"

	"hash/crc32"

	"github.com/Vealcoo/go-pkg/conversion"
)

func StringHash(in string) string {
	h := crc32.NewIEEE()
	h.Write([]byte(in))

	return conversion.Decimal2Other(int64(h.Sum32()), 62)
}

func StringMD5(in string) string {
	has := md5.Sum([]byte(in))
	md5str := fmt.Sprintf("%x", has)

	return md5str
}
