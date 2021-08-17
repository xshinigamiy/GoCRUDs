package request

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	UpdatedBy string `json:"updatedBy"`
	// UpdatedAt string `json:"updatedAt"`
	CreatedBy string `json:"createdBy"`
	// CreatedAt  string `json:"createdAt"`
	Profession string `json:"profession"`
}

// type Update_User struct {

// }
