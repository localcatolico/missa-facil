package entity

import "errors"

var (
	ErrInvalidEntity  = errors.New("invalid-entity")
	ErrNotFoundEntity = errors.New("not-found-entity")
)
