package model

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type AddUser struct {
	Name string `json:"name" binding:"required"`
	Age  uint8  `json:"age" binding:"required"`
}

type UpdateUser struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type UpdateUserSwagger struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}
