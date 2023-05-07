package dto

type PaginationDto struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
