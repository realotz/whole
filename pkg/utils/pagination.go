package utils

type Query interface {
	Offset(offset int)
	Limit(limit int)
}

func GetOfficeLimit(page, pageSize *uint32) (int, int) {
	if *page <= 0 {
		*page = 1
	}
	if *pageSize <= 0 {
		*pageSize = 10
	}
	if *pageSize > 200 {
		*pageSize = 200
	}
	return int((*page - 1) * *pageSize), int(*pageSize)
}
