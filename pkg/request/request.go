package request

type User struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Age        int    `json:"age"`
	UpdatedBy  string `json:"updatedBy"`
	UpdatedAt  string `json:"updatedAt"`
	CreatedBy  string `json:"createdAt"`
	CreatedAt  string `json:"createdBy"`
	Profession string `json:"profession"`
}
