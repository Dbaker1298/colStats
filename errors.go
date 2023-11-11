package main

import (
	"errors"
)

var (
	ErrNotNumber        = errors.New("Data is not numberic")
	ErrInvalidColumn    = errors.New("Invalid column number")
	ErrNoFiles          = errors.New("No files to process")
	ErrInvalidOperation = errors.New("Invalid operation")
)
