package endpoint

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/labstack/gommon/log"
)

func ResponseErrorHandler(re *errors.Response) {
	log.Errorf("ResponseErrorHandler: %s", re.Error.Error())
}
