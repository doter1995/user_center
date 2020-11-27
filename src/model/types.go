package model

type Pagination struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
	Total    int `json:"total"`
	Data     []interface{}
}

type UserPagination struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
	Total    int `json:"total"`
	Data     []User
}

type AppPagination struct {
	PageSize int `json:"pageSize"`
	Page     int `json:"page"`
	Total    int `json:"total"`
	Data     []App
}
