package check

import (
	"golang.org/x/exp/constraints"
)

// Intersect : 交集
func Intersect[T constraints.Ordered](c1, c2 []T) []T {
	c1Map := map[T]bool{}

	for _, v := range c1 {
		c1Map[v] = true
	}

	res := []T{}

	for _, v := range c2 {
		if c1Map[v] {
			res = append(res, v)
		}
	}

	return res
}

// Difference : 差集（c2 - c1）
func Difference[T constraints.Ordered](c1, c2 []T) []T {
	c1Map := map[T]bool{}

	for _, v := range c1 {
		c1Map[v] = true
	}

	res := []T{}

	for _, v := range c2 {
		if c1Map[v] {
			continue
		}
		res = append(res, v)
	}

	return res
}

// Union : 聯集
func Union[T constraints.Ordered](c1, c2 []T) []T {
	c1Map := map[T]bool{}

	for _, v := range c1 {
		c1Map[v] = true
	}

	res := []T{}
	res = append(res, c1...)

	for _, v := range c2 {
		if c1Map[v] {
			continue
		}
		res = append(res, v)
	}

	return res
}

// IsExist : 檢查 set 中是否存在 e
func IsExist[T constraints.Ordered](set []T, e T) bool {
	for _, v := range set {
		if v == e {
			return true
		}
	}

	return false
}

// DeleteElements : 移除 set 中 e 的元素
func DeleteElements[T constraints.Ordered](set []T, e []T) []T {
	res := []T{}

	eMap := map[T]bool{}
	for _, v := range e {
		eMap[v] = true
	}

	for _, v := range set {
		if !eMap[v] {
			res = append(res, v)
		}
	}

	return res
}
