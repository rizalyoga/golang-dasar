package request

type BookRequest struct {
	Title       string `json:"title" validate:"required,min=3"`
	AuthorID    uint   `json:"author_id" validate:"required"`
	Description string `json:"description" validate:"required,min=3"`
}
