package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("urls not found")
	ErrUrlExists   = errors.New("urls already exists")
)
