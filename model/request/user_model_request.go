package model_request

type CreateUserRequest struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}

type UpdateUserRequest struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}