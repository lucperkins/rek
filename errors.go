package rek

import "errors"

// The error thrown when you attempt to set the HTTP request body more than once
var ErrRequestBodySetMultipleTimes = errors.New("you have set the request body more than once")
