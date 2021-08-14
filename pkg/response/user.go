package response

type UserResponse struct {
	Id       int    `json:"userId"`
	FullName string `json:"fullName"`
}
