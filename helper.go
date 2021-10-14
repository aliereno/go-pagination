package pagination

func calculateTotalPage(totalSize, pageSize int) (totalPage int) {
	totalPage = totalSize / pageSize
	if totalSize%pageSize > 0 {
		totalPage++
	}
	return
}
