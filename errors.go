package rek

import "errors"

var ErrRequestBodySetMultipleTimes = errors.New("you have set the request body more than once")
