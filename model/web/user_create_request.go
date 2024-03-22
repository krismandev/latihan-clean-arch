package web

type UserCreateRequest struct {
	Username  string `validate:"required,min=6" json:"username"`
	Password  string `validate:"required,min=8" json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
