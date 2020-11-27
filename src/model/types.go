package model

type Pagination struct {
	PageSize int           `json:"pageSize"`
	Page     int           `json:"page"`
	Total    int           `json:"total"`
	Data     []interface{} `json:"data"`
}

type UserPagination struct {
	PageSize int    `json:"pageSize"`
	Page     int    `json:"page"`
	Total    int    `json:"total"`
	Data     []User `json:"data"`
}

type AppPagination struct {
	PageSize int   `json:"pageSize"`
	Page     int   `json:"page"`
	Total    int   `json:"total"`
	Data     []App `json:"data"`
}
