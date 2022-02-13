package model_response

type CreateCustomerResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserRole  string `json:"user_role"`
}

type ListCustomerResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserRole  string `json:"user_role"`
}

type UpdateCustomerResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserRole  string `json:"user_role"`
}