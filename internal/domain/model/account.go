package model

type Account struct {
	Id       uint64 `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"pass_hash"`
}

type RegisterBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	AccountId uint64 `json:"account_id"`
}
