package conversion

var chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Decimal2Other(num int64, base int64) string {
	if base < 2 || base > 62 {
		return ""
	}

	return encode(num, base)
}

func encode(num int64, base int64) string {
	bytes := []byte{}
	// FIXME: if num <= 0
	for num > 0 {
		bytes = append(bytes, chars[num%base])
		num = num / base
	}
	reverse(bytes)
	return string(bytes)
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
