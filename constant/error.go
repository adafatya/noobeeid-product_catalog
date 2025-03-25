package constant

type Error struct {
	HttpStatus   int
	Message      string
	ErrorMessage string
	ErrorCode    int
}

var (
	ErrRepositoryError = &Error{
		HttpStatus:   500,
		Message:      "internal server error",
		ErrorMessage: "error repository",
		ErrorCode:    50001,
	}
	ErrUnknownError = &Error{
		HttpStatus:   500,
		Message:      "internal server error",
		ErrorMessage: "unknown error",
		ErrorCode:    99999,
	}

	// auth error
	ErrRequiredEmail = &Error{
		HttpStatus:   400,
		Message:      "bad request",
		ErrorMessage: "email required",
		ErrorCode:    40001,
	}
	ErrInvalidEmail = &Error{
		HttpStatus:   400,
		Message:      "bad request",
		ErrorMessage: "invalid email",
		ErrorCode:    40002,
	}
	ErrRequiredPassword = &Error{
		HttpStatus:   400,
		Message:      "bad request",
		ErrorMessage: "password required",
		ErrorCode:    40003,
	}
	ErrInvalidPassword = &Error{
		HttpStatus:   400,
		Message:      "bad request",
		ErrorMessage: "invalid password",
		ErrorCode:    40004,
	}
	ErrDuplicateEmail = &Error{
		HttpStatus:   409,
		Message:      "bad request",
		ErrorMessage: "email already used",
		ErrorCode:    40901,
	}
)
