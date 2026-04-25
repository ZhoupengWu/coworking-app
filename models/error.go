package models

const (
	InvalidCredentialsErr = "User not signed up or bad credentials"
	EmailAlreadyInUseErr  = "Email already signed up"
	DbErr                 = "Generic db error"
	ObjectNotFoundErr     = "Object not found with the requested id"
)

type CoworkingErr struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (c CoworkingErr) Error() string {
	return c.Message
}
