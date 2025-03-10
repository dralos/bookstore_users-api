package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/dralos/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error processing database request")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("duplicated data")
	}
	return errors.NewInternalServerError(fmt.Sprintf("error processing request %v", sqlErr))
}
