package check

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CheckSuite struct {
	suite.Suite
}

func TestCheckSuite(t *testing.T) {

	suite.Run(t, &CheckSuite{})
}

func (c *CheckSuite) TestIntersectWithString() {
	c1 := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	c2 := []string{"a", "c", "d", "e", "z"}

	expect := []string{"a", "c", "d", "e"}
	resp := Intersect(c1, c2)

	c.Equal(expect, resp)
}

func (c *CheckSuite) TestDifferenceWithString() {
	c1 := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	c2 := []string{"a", "c", "d", "e", "z"}

	expect := []string{"z"}
	resp := Difference(c1, c2)

	c.Equal(expect, resp)
}

func (c *CheckSuite) TestDifferenceWithInt() {
	c1 := []int{1, 2, 3, 4, 5, 6, 7}
	c2 := []int{1, 2, 3, 4, 10, 11, 12}

	expect := []int{10, 11, 12}
	resp := Difference(c1, c2)

	c.Equal(expect, resp)
}

func (c *CheckSuite) TestUnionWithString() {
	c1 := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	c2 := []string{"a", "c", "d", "e", "z"}

	expect := []string{"a", "b", "c", "d", "e", "f", "g", "h", "z"}
	resp := Union(c1, c2)

	c.Equal(expect, resp)
}

func (c *CheckSuite) TestIsExist() {
	set := []string{"A", "B", "C", "D"}
	c.True(IsExist(set, "C"))
	c.False(IsExist(set, "Z"))
}

func (c *CheckSuite) TestDeleteElements() {
	set := []string{"A", "B", "C", "D", "E", "F"}
	es := []string{"C", "F"}

	res := DeleteElements(set, es)
	c.NotContains(res, es)
}
