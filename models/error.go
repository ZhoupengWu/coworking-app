package models

const (
	InvalidCredentialsErr = "User not signed up or bad credentials"
	EmailAlreadyInUseErr  = "Email already signed up"
	DbErr                 = "Generic db error"
	ObjectNotFoundErr     = "Object not found with the requested id"
	ValidationErr         = "Body validation"
	TokenGenerationErr    = "Failure in generating the JWT token"
	DateWrongFormatErr    = "Date has a wrong format. Expected YYYY-MM-DD"
	MissingTokenErr       = "The JWT token is missing"
	TokenNotValidErr      = "The JWT token is not valid"
)

type CoworkingErr struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (c CoworkingErr) Error() string {
	return c.Message
}
