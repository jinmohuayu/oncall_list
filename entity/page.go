package entity

// Page 分页参数
type Page struct {
	CurrentPageIndex int         // 当前页码(首页为1)
	RecordsPerPage   int         // 每页记录数
	TotalRecords     int         // 总记录数
	Records          interface{} // 记录队列
}

// TotalPages 总的页数
func (p Page) TotalPages() int {
	if p.RecordsPerPage == 0 {
		return 0
	}

	return (p.TotalRecords + p.RecordsPerPage - 1) / p.RecordsPerPage
}
