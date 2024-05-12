package data

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Model struct {
	Movies MovieModel
}
