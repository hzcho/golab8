package model

type GetUserFilter struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

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
