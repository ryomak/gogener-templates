package merr

import (
	"database/sql"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

/*
1xxx...validation
2xxx...auth
3xxx...DB
9999...unexpectd
*/
const (
	CodeValidator            = "1000"
	CodeAuthFailed           = "2000"
	CodeAuthPhoneUnregisterd = "2001"
	CodeDBNotFound           = "3000"
	CodeDBMaintenance        = "3001"
	CodeUnexpected           = "9999"
)

var (
	ErrAuthPhoneUnregisterd = &Error{
		"firebae phoneNumber unregisterd",
		CodeAuthPhoneUnregisterd,
	}
	ErrMaintenance = &Error{
		"maintenance now.",
		CodeDBMaintenance,
	}
)

type Error struct {
	Msg  string
	Code string
}

func CreateError(msg, code string) *Error {
	return &Error{
		msg,
		code,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("msg=%s", e.Msg)
}

func IsMyError(err error) bool {
	if err == nil {
		return false
	}
	switch err.(type) {
	case *Error:
		return true
	default:
		return false

	}
}

func IsErrNoRows(err error) bool {
	return err == sql.ErrNoRows
}

func IsErrValidator(err error) bool {
	_, ok := err.(validator.ValidationErrors)
	return ok
}
