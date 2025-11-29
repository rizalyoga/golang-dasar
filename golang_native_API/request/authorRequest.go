package request

type AuthorRequest struct {
	Name   string `json:"name" validate:"required,min=3"`
	Gender string `json:"gender" validate:"required,oneof=M F"`
	Email  string `json:"email" validate:"required,email"`
	Age    int    `json:"age" validate:"required,numeric"`
}
