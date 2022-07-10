package random

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RandomSuite struct {
	suite.Suite
	testTimes int
}

func TestRandomSuite(t *testing.T) {

	suite.Run(t, &RandomSuite{
		testTimes: 1000,
	})
}

func (t *RandomSuite) TestGetNumUint32() {
	min := uint32(5)
	max := uint32(10)

	for i := 1; i <= t.testTimes; i++ {
		num := GetNum(min, max)
		t.GreaterOrEqual(num, int64(min))
		t.LessOrEqual(num, int64(max))
	}
}

func (t *RandomSuite) TestGetNumInt64() {
	min := int64(1)
	max := int64(10)

	for i := 1; i <= t.testTimes; i++ {
		num := GetNum(min, max)
		t.GreaterOrEqual(num, min)
		t.LessOrEqual(num, max)
	}
}

func (t *RandomSuite) TestGetNumByCount() {
	min, max, count := 1, 10, 100

	numSlice := GetNumByCount(min, max, count)

	t.Len(numSlice, int(count))

	for _, num := range numSlice {
		t.GreaterOrEqual(num, int64(min))
		t.LessOrEqual(num, int64(max))
	}
}

func (t *RandomSuite) TestGetNumByCountNoRepeat() {
	min, max, count := 1, 10, 8

	numSlice := GetNumByCountNoRepeat(min, max, count)

	t.Len(numSlice, int(count))

	m := map[int64]struct{}{}

	for _, num := range numSlice {
		t.NotContains(m, num)
		m[num] = struct{}{}

		t.GreaterOrEqual(num, int64(min))
		t.LessOrEqual(num, int64(max))
	}
}
