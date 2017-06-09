package depdemolib

import (
	"github.com/pkg/errors"
)

func Msg() string {
	msg := "hello world"
	if false {
		panic(errors.New(msg))
	}
	return msg
}
