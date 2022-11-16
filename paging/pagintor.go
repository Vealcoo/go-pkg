package paging

// Paginator : 分頁器
func Paginator[T interface{}](in []T, page, limit int) []T {
	if limit == 0 || page == 0 {
		return nil
	}

	start := (page - 1) * limit
	end := start + limit

	pageCount := len(in) / limit
	if pageCount == 0 || len(in)%limit > 0 {
		pageCount++
	}

	if page > pageCount {
		return nil
	}

	if end > len(in) {
		end = len(in)
	}

	return in[start:end]
}
