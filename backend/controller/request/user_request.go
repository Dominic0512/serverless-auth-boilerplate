package request

type ManipulateRequest struct {
	ID string `uri:"id" validate:"required"`
}

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type PartialUpdateUserRequest struct {
	Name string `json:"name" validate:""`
}
