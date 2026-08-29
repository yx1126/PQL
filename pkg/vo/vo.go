package vo

type BaseVo struct {
	Id int `json:"id" form:"id" binding:"required"`
}

type Result[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}
