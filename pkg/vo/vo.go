package vo

type BaseVo struct {
	Id int `json:"id" form:"id" binding:"required"`
}

type Result[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type ResultPaging[T any] struct {
	Total int `json:"total"`
	Data  []T `json:"data"`
}
