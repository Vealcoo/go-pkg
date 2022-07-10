package random

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetNum : 範圍中取得一個隨機數 可帶入 int / uint 的各種型態
func GetNum[T constraints.Integer](min, max T) int64 {
	return rand.Int63n(int64(max-min)+1) + int64(min)
}

// GetNumByCount : 取得一定數量的隨機數
func GetNumByCount[T constraints.Integer](min, max, count T) []int64 {
	res := []int64{}

	for count > 0 {
		res = append(res, GetNum(min, max))
		count--
	}

	return res
}

// GetNumByCountNoRepeat : 取得一定數量的不重複隨機數
func GetNumByCountNoRepeat[T constraints.Integer](min, max, count T) []int64 {
	// (max - min + 1) 必須 >= count ， 否則無法取得不重複隨機數
	if (max - min + 1) < count {
		return nil
	}

	res := []int64{}
	m := map[int64]struct{}{}

	for count > 0 {
		n := GetNum(min, max)

		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
			res = append(res, n)
			count--
		}
	}

	return res
}
