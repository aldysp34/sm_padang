package apperr

import (
	"errors"
	"net/http"
)

type CustomError struct {
	message string
	Code    int `json:"code"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		message: message,
	}
}

func (c *CustomError) Error() string {
	return c.message
}

func (c *CustomError) ConvertToErrorResponse() ErrorResponse {
	return ErrorResponse{
		Message: c.message,
	}
}

var (
	ErrDatabaseConnection = errors.New("cannot make connection to mongo server")
	ErrSetCookie          = NewCustomError(http.StatusUnauthorized, "failed to set cookie")
	ErrInsertNewUser      = NewCustomError(http.StatusInternalServerError, "failed when create new user")
	ErrCookieNotFound     = NewCustomError(http.StatusUnauthorized, "cookie not found")
	ErrInsertPerfume      = NewCustomError(http.StatusInternalServerError, "insert perfume query error")
	ErrFindPerfumes       = NewCustomError(http.StatusInternalServerError, "find perfume query error")
	ErrInvalidID          = NewCustomError(http.StatusBadRequest, "id not found")
	ErrFindPerfume        = NewCustomError(http.StatusInternalServerError, "find perfume query error")
	ErrCreateProductType  = NewCustomError(http.StatusInternalServerError, "insert perfume type query error")
	ErrFindPerfumeType    = NewCustomError(http.StatusInternalServerError, "find perfume type query error")
	ErrInsertImages       = NewCustomError(http.StatusInternalServerError, "insert image query error")
	ErrFindImages         = NewCustomError(http.StatusInternalServerError, "find images query error")
	ErrFailedTransaction  = NewCustomError(http.StatusInternalServerError, "transaction query error")
	ErrInsertArticle      = NewCustomError(http.StatusInternalServerError, "insert Article query error")
	ErrFindArticles       = NewCustomError(http.StatusInternalServerError, "find Article query error")
	ErrScanUser           = NewCustomError(http.StatusInternalServerError, "failed to scan to user")
	ErrFindUsersQuery     = NewCustomError(http.StatusInternalServerError, "find user query error")
	ErrFindUserByIdQuery  = NewCustomError(http.StatusInternalServerError, "find user by id query error")
	ErrFindUserByEmail    = NewCustomError(http.StatusInternalServerError, "find user by email query error")
	ErrNewUserQuery       = NewCustomError(http.StatusInternalServerError, "new user query error")
	ErrUserNotFound       = NewCustomError(http.StatusBadRequest, "user not found")
	ErrEmailAlreadyUsed   = NewCustomError(http.StatusBadRequest, "email already used")
	ErrWrongCredentials   = NewCustomError(http.StatusUnauthorized, "wrong password or email/username")
	ErrWrongEmail         = NewCustomError(http.StatusUnauthorized, "email not registered")
	ErrUpdateUserDetails  = NewCustomError(http.StatusInternalServerError, "user info is not updated")
	ErrExpiredToken       = NewCustomError(http.StatusBadRequest, "Token is invalid or expired")
	ErrEditProduct        = NewCustomError(http.StatusBadRequest, "failed to edit product")
	ErrDeleteProduct      = NewCustomError(http.StatusBadRequest, "failed to Delete product")
	ErrEditArticle        = NewCustomError(http.StatusBadRequest, "failed to edit Article")
	ErrDeleteArticle      = NewCustomError(http.StatusBadRequest, "failed to Delete Article")

	ErrNewVerifyRequest    = NewCustomError(http.StatusBadRequest, "failed to create forgot password request")
	ErrDeleteVerifyRequest = NewCustomError(http.StatusInternalServerError, "query delete verification request")

	ErrEmailUnregistered = NewCustomError(http.StatusInternalServerError, "user email never registered")

	ErrRecordNotFound       = NewCustomError(http.StatusBadRequest, "record not found")
	ErrRecordDuplicateFound = NewCustomError(http.StatusBadRequest, "record duplicate found")
	ErrDatabaseQuery        = NewCustomError(http.StatusInternalServerError, "database query error")

	ErrGenerateHashPassword = NewCustomError(http.StatusInternalServerError, "couldn't generate hash password")
	ErrGenerateJWTToken     = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")

	ErrTxCommit = NewCustomError(http.StatusInternalServerError, "commit transaction error")

	ErrInvalidBody         = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrUnauthorize         = NewCustomError(http.StatusUnauthorized, "unauthorized")
	ErrBearerTokenNotExist = NewCustomError(http.StatusUnauthorized, "bearer token not exist")

	ErrNewResetToken         = NewCustomError(http.StatusInternalServerError, "failed to create reset token")
	ErrInvalidateResetToken  = NewCustomError(http.StatusInternalServerError, "failed to update token valid status")
	ErrFindTokenData         = NewCustomError(http.StatusInternalServerError, "error on find password reset token data")
	ErrTokenDataNotFound     = NewCustomError(http.StatusBadRequest, "token data not found")
	ErrTokenIsInvalid        = NewCustomError(http.StatusInternalServerError, "token is invalid")
	ErrTokenIsUsed           = NewCustomError(http.StatusInternalServerError, "token is already used")
	ErrResetPasswordNotFound = NewCustomError(http.StatusInternalServerError, "token for reset password not found")
	ErrInvalidSignature      = NewCustomError(http.StatusBadRequest, "invalid signature")
	ErrBearerTokenInvalid    = NewCustomError(http.StatusBadRequest, "invalid bearer token")
)
