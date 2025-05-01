package verify

type VerifyCreateRequest struct {
	Email string `json:"email" validate:"required,email"`
}
