package property

import "errors"

var (
	ErrorAlreadyResident = errors.New("already resident")
	ErrorHasResidence    = errors.New("already has residence")
)
