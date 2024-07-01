package models

import "errors"

var (
	ErrEmailTaken    = errors.New("models: email address is already in use")
	ErrWrongPassword = errors.New("models: wrong password given as input")
	ErrUserNotExist  = errors.New("models: user with that email address doesn't exist")
	ErrNotFound      = errors.New("models: resource could not be found")
)
