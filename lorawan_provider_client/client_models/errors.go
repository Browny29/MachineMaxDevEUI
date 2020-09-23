package client_models

import "errors"

var (
	ErrDevEUIAlreadyExists = errors.New("this DevEUI already exists on the server, please try again")
)

