package exception

import "errors"

const (
	StatusInternalServer = "Internal Server Error"
	StatusBadRequest     = "Bad Request"
	StatusSuccess        = "Success"
	StatusUnauthorized   = "Unauthorized"
)

var (
	ErrNotFound                 = errors.New("data not found")
	ErrUsernameAlreadyExist     = errors.New("username already exist")
	ErrEmailAlreadyExist        = errors.New("email already exist")
	ErrPhoneNumberAlreadyExist  = errors.New("phone number already exist")
	ErrCardNumberAlreadyExist   = errors.New("card number already exist")
	ErrCvvAlreadyExist          = errors.New("cvv already exist")
	ErrMinimalTransaction       = errors.New("amount must be greater than 10000")
	ErrNotEnoughBalance         = errors.New("not enough balance")
	ErrFailedCreate             = errors.New("failed to create data")
	ErrFailedCreateToken        = errors.New("failed to create token")
	ErrInvalidParseToken        = errors.New("invalid to parse token")
	ErrInvalidTokenMapclaims    = errors.New("invalid token mapclaims ")
	ErrInvalidTokenStringMethod = errors.New("invalid token string method")
	ErrInvalidExtension         = errors.New("extention is not allowed")
	ErrInvalidUsernamePassword  = errors.New("invalid username password")
	ErrTokenNotProvided         = errors.New("token not provided")
	ErrFailedUpdate             = errors.New("failed to update data")
	ErrFailedUpload             = errors.New("failed to upload data")
	ErrFailedDelete             = errors.New("failed to delete data")
	ErrTitleAlreadyExist        = errors.New("title already exist")
	ErrInvalidPage              = errors.New("invalid page")
	ErrInvalidPerPage           = errors.New("invalid per page")
)
