package hash

import "hash/crc32"

func StringHash(in string) uint32 {
	h := crc32.NewIEEE()
	h.Write([]byte(in))

	return h.Sum32()
}
