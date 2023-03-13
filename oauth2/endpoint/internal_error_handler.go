package endpoint

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/gommon/log"
)

func InternalErrorHandler(err error) (re *errors.Response) {
	log.Errorf("InternalErrorHandler: %s", err.Error())
	return re
}
