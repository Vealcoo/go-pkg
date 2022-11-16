package paging

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PagingSuite struct {
	suite.Suite
}

func TestPagingSuite(t *testing.T) {

	suite.Run(t, &PagingSuite{})
}

func (t *PagingSuite) TestPaginator() {
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	page := 2
	limit := 10

	result := Paginator(testData, page, limit)

	expectData := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	t.Equal(expectData, result)
}
